//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

import "openzeppelin-solidity/contracts/token/ERC20/IERC20.sol";
import "hardhat/console.sol";

contract BidWar {
    address private commissionContract;
    uint256 private commissionPercentage;
    uint256 private duration;
    uint256 private increment;
    uint256 private finalTimestamp;
    address private topBidder;
    uint256 private totalBidAmount = 0;
    uint256 private previousBid = 0;
    mapping(string => uint256) private allBids;

    constructor(uint256 _duration, uint256 _increment, uint256 _commissionPercentage) {
        duration = _duration;
        increment = _increment;
        commissionContract = msg.sender;
        commissionPercentage = _commissionPercentage;
    }

    event NewBid(address bidder, uint256 amount, uint256 finalTimestamp);
    event Timeout(address bidder, uint256 amount);

    function newBid(uint256 amount, address tokenAddress) external {
        require(amount > previousBid, "LowerBidAmount");

        IERC20 token = IERC20(tokenAddress);
        uint256 allowance = token.allowance(msg.sender, address(this));
        require(allowance >= amount, "IncorrectAllowance");

        uint256 commission = (amount * commissionPercentage) / 100;
        bool success = token.transferFrom(
            msg.sender,
            commissionContract,
            commission
        );
        require(success, "TransferFailure");


        bool transferSuccess = token.transferFrom(
            msg.sender,
            address(this),
            amount - commission
        );
        require(transferSuccess, "TransferFailure");
        // allBids[abi.encodePacked(msg.sender)] += amount;
        previousBid = amount;
        topBidder = msg.sender;
        if (totalBidAmount == 0) {
            finalTimestamp = block.timestamp + duration;
        } else {
            finalTimestamp += increment;
        }
        totalBidAmount += amount - commission;
        emit NewBid(msg.sender, amount, finalTimestamp);
    }

    function finalise(address tokenAdress) public {
        // console.log(
        //     "Final time: %s Current time: %s",
        //     finalTimestamp,
        //     block.timestamp
        // );
        require(block.timestamp >= finalTimestamp, "BidTimeLeft");

        IERC20 token = IERC20(tokenAdress);
        token.transfer(topBidder, totalBidAmount);
        emit Timeout(topBidder, totalBidAmount);
        // reset parameters
        totalBidAmount = 0;
        finalTimestamp = 0;
        topBidder = address(0);
        previousBid = 0;
    }
}
