// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0 <0.9.0;

contract Voting {
    mapping(address => bool) private voted;
    mapping(address => uint256) private votes;
    address owner;
    bool isOver;

    constructor() {
      	owner = msg.sender;
        isOver = false;
    }

    function vote(uint broj) public notVoted votingNotOver {
        voted[msg.sender] = true;
        votes[msg.sender] = broj;
    }

    function closeVoting() public onlyByOwner {
        isOver = true;
    }

    modifier notVoted {
        require(voted[msg.sender], "Sender already voted.");
		_;
    }

    modifier votingNotOver {
        require(isOver == false, "Voting is over.");
		_;
    }

    modifier onlyByOwner {
        require(msg.sender == owner, "Sender not authorized.");
        _;
    }
}
