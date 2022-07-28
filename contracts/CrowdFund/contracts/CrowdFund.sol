// SPDX-License-Identifier: MIT
// ERC20众筹合约
pragma solidity ^0.8.10;

import "contracts/openzeppelin-contracts/contracts/token/ERC20/IERC20.sol";

contract CrowdFund {
    event Launch(
        uint id,
        address indexed creator,
        uint goal,
        uint32 startedAt,
        uint32 endAt
    );

    event Cancel(uint _id);
    event Pledge(uint indexed _id, address indexed caller, uint _amount);
    event UnPledge(uint indexed _id, address indexed caller, uint _amount);
    event Claim(uint _id);
    event Refund(uint indexed id, address indexed caller, uint amount);

    struct Campaign {
        address creator;
        uint goal;
        uint pledged;
        uint32 startedAt;
        uint32 endAt;
        bool claimed;
    }

    IERC20 public immutable token;

    uint public count;
    mapping(uint => Campaign) public campaigns;
    mapping(uint => mapping(address => uint)) public pledgedAmount;

    constructor (address _token) {
        token = IERC20(_token);
    }

    // 创建众筹
    function launch(
        uint _goal,
        uint32 _startAt,
        uint32 _endAt
    ) external {
        require(_startAt >= block.timestamp, "start at < now");
        require(_endAt >= _startAt, "ent at < start at");
        require(_endAt <= block.timestamp + 90 days, "end at > max duration");

        count += 1;

        campaigns[count] = Campaign({
            creator : msg.sender,
            goal : _goal,
            pledged : 0,
            startedAt : _startAt,
            endAt: _endAt,
            claimed : false
        });
        emit Launch(count, msg.sender, _goal, _startAt, _endAt);
    }

    // 取消众筹
    function cancel(uint _id) external {
        Campaign memory campaign = campaigns[_id];
        require(msg.sender == campaign.creator, "not creator");
        require(block.timestamp >= campaign.startedAt, "started");
        delete campaigns[_id];
        emit Cancel(_id);
    }

    // 参与众筹
    function pledge(uint _id, uint _amount) external {
        Campaign storage campaign = campaigns[_id];
        require(block.timestamp >= campaign.startedAt, "not started");
        require(block.timestamp <= campaign.endAt, "ended");

        campaign.pledged += _amount;
        pledgedAmount[_id][msg.sender] += _amount;
        token.transferFrom(msg.sender, address(this), _amount);
        emit Pledge(_id, msg.sender, _amount);
    }

    // 退出众筹
    function unPledge(uint _id, uint _amount) external {
        Campaign storage campaign = campaigns[_id];
        require(block.timestamp <= campaign.endAt, "ended");
        campaign.pledged -= _amount;
        pledgedAmount[_id][msg.sender] -= _amount;
        token.transfer(msg.sender, _amount);

        emit UnPledge(_id, msg.sender, _amount);
    }

    // 众筹成功取款
    function claim(uint _id) external {
        Campaign storage campaign = campaigns[_id];
        require(msg.sender == campaign.creator, "not creator");
        require(block.timestamp > campaign.endAt, "not ended");
        require(campaign.pledged >= campaign.goal, "pledged < goal");
        require(!campaign.claimed, "claimed");

        campaign.claimed = true;
        token.transfer(msg.sender, campaign.pledged);
        emit Claim(_id);
    }

    // 众筹失败退款
    function refund(uint _id) external {
        Campaign storage campaign = campaigns[_id];
        require(block.timestamp > campaign.endAt, "not ended");
        require(campaign.pledged < campaign.goal, "pledged < goal");
        uint balance = pledgedAmount[_id][msg.sender];
        pledgedAmount[_id][msg.sender] = 0;
        token.transfer(msg.sender, balance);
    }
}

