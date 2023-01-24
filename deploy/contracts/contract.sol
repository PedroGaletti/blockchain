// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.4;

contract Contract {
    // state variable to store secret number
    uint256 secretNumber;

    // set the secret number on deploy
    constructor() payable {
        secretNumber = 33;
    }

    // declare the event with the payload to send
    event LotteryEvent(bool isWinner, address indexed player);

    //
    function guessNumber(uint256 _number) public payable {
        if (_number != secretNumber) {
            // EVENT
            emit LotteryEvent(false, msg.sender);
        } else {
            // EVENT
            emit LotteryEvent(true, msg.sender);
        }
    }
}
