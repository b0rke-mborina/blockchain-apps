// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0 <0.9.0;

contract SimpleBank {
    // balances mapping
    mapping(address => uint256) private balances;

    // deposits money to contract
    function deposit() external payable {
        balances[msg.sender] += msg.value;
    }

    // withdraws sender's money from contract
    function withdraw(uint256 withdrawAmount) external {
        require(balances[msg.sender] >= withdrawAmount, "ERROR: Insufficient balance");
        payable(msg.sender).transfer(withdrawAmount);
        balances[msg.sender] -= withdrawAmount;
    }

    // checks sender's balance deposited on the contract
    function balance() external view returns (uint256) {
        return balances[msg.sender];
    }

    // returns the full balance of the contract
    function depositsBalance() external view returns (uint256) {
        return address(this).balance;
    }
}