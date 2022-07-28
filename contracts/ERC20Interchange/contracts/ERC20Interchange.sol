// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../../openzeppelin-contracts-upgradeable/contracts/token/ERC20/IERC20Upgradeable.sol";
import "../../openzeppelin-contracts-upgradeable/contracts/proxy/utils/Initializable.sol";
import "../../openzeppelin-contracts-upgradeable/contracts/utils/CountersUpgradeable.sol";

contract ERC20Interchange is Initializable{
    event Create(uint256 indexed interchangeId, address from, address to, uint256 from_amount, uint256 to_amount, uint8 interchange_type);
    event InterchangeSuccess(uint256 interchangeId);

    using CountersUpgradeable for CountersUpgradeable.Counter;
    CountersUpgradeable.Counter private interchangeCounter;

    struct Interchange {
        address from_address;
        address to_address;
        uint256 from_amount;
        uint256 to_amount;
        uint8 interchange_type; // 0: token_one -> token_two, 1: token_two -> token_one
    }

    IERC20Upgradeable public token_one;
    IERC20Upgradeable public token_two;

    // Mapping from owner to list of owned Interchange IDs
    mapping(address => mapping(uint256 => uint256)) private _fromInterchanges;

    // Mapping to owner to list of owned Interchange IDs
    mapping(address => mapping(uint256 => uint256)) private _toInterchanges;

    // Mapping from pending count
    mapping(address => uint256) private _fromInterchangeCount;

    // Mapping to pending count
    mapping(address => uint256) private _toInterchangeCount;

    // Mapping from interchange ID to index of the owner interchanges list
    mapping(uint256 => uint256) private _fromInterchangesIndex;

    // Mapping from interchange ID to index of the owner interchanges list
    mapping(uint256 => uint256) private _toInterchangesIndex;

    // Mapping from interchange id to position in the allInterchanges array
    mapping(uint256 => Interchange) private allInterchanges;

    function initialize(address _token_one, address _token_two) initializer public {
        token_one = IERC20Upgradeable(_token_one);
        token_two = IERC20Upgradeable(_token_two);
    }

    function create(address to_address, uint256 from_amount, uint256 to_amount, uint8 interchange_type) public returns (uint256 interchangeId) {
        require(interchange_type < 2, "interchange_type error");
        if (interchange_type == 0 ) {
            require(token_one.balanceOf(msg.sender) > from_amount, "token_one balance insufficient balance");
        } else {
            require(token_two.balanceOf(msg.sender) > from_amount, "token_two balance insufficient balance");
        }
        interchangeCounter.increment();
        interchangeId = interchangeCounter.current();
        allInterchanges[interchangeId] = Interchange({
            from_address: msg.sender,
            to_address: to_address,
            from_amount: from_amount,
            to_amount: to_amount,
            interchange_type: interchange_type
        });

        if (interchange_type == 0 ) {
            token_one.transferFrom(msg.sender, address(this), from_amount);
        } else {
            token_two.transferFrom(msg.sender, address(this), from_amount);
        }

        interchangeCreate(interchangeId, to_address);

        emit Create(interchangeId, msg.sender, to_address, from_amount, to_amount, interchange_type);
    }

    function interchangeCreate(uint256 interchangeId, address to_address) internal{
        _fromInterchanges[msg.sender][_fromInterchangeCount[msg.sender]] = interchangeId;
        _toInterchanges[to_address][_toInterchangeCount[to_address]] = interchangeId;
        _fromInterchangesIndex[interchangeId] = _fromInterchangeCount[msg.sender];
        _toInterchangesIndex[interchangeId] = _toInterchangeCount[to_address];
        _fromInterchangeCount[msg.sender] += 1;
        _toInterchangeCount[to_address] += 1;
    }

    function interchangeCompleted(uint256 interchangeId, address from_address) internal{
        _fromInterchangeCount[from_address] -= 1;
        _toInterchangeCount[msg.sender] -= 1;
        delete _fromInterchangesIndex[interchangeId];
        delete _toInterchangesIndex[interchangeId];
        delete _toInterchanges[msg.sender][_toInterchangesIndex[interchangeId]];
        delete _fromInterchanges[from_address][_fromInterchangesIndex[interchangeId]];
        delete allInterchanges[interchangeId];
    }

    function fromList() public view returns(uint256[] memory) {
        uint256[] memory res = new uint256[](_fromInterchangeCount[msg.sender]);
        for (uint256 i = 0; i < _fromInterchangeCount[msg.sender]; i ++ ) {
            uint256 interchangeId = _fromInterchanges[msg.sender][i];
            res[i] = interchangeId;
        }
        return res;
    }

    function toList() public view returns(uint256[] memory) {
        uint256[] memory res = new uint256[](_toInterchangeCount[msg.sender]);
        for (uint256 i = 0; i < _toInterchangeCount[msg.sender]; i ++ ) {
            uint256 interchangeId = _toInterchanges[msg.sender][i];
            res[i] = interchangeId;
        }
        return res;
    }

    function searchByInterchangeId(uint256 interchangeId) public view returns(Interchange memory){
         return allInterchanges[interchangeId];
    }

    function interchange(uint256 interchangeId, uint256 amount) public {
        require(interchangeId <= interchangeCounter.current(), "interchangeId not exist");
        Interchange storage data = allInterchanges[interchangeId];
        require(data.from_address != address(0), "interchangeId completed");
        require(msg.sender == data.to_address, "interchangeId error");
        require(amount >= data.to_amount, "amount < expect_amount");
        if (data.interchange_type == 0) {
            require(token_two.balanceOf(msg.sender) >= amount, "token_two balance insufficient balance");
            token_one.transfer(msg.sender, data.from_amount);
            token_two.transferFrom(msg.sender, data.from_address, amount);
        } else {
            require(token_one.balanceOf(msg.sender) >= amount, "token_one balance insufficient balance");
            token_two.transfer(msg.sender, data.from_amount);
            token_one.transferFrom(msg.sender, data.from_address, amount);
        }
        interchangeCompleted(interchangeId, data.from_address);
        emit InterchangeSuccess(interchangeId);
    }
}
