// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0 <0.9.0;

contract TemperatureControl {
    address public owner;
    mapping(address => bool) public authorizedAddresses;
    uint public temperature;

    // remebers owner of contract
    constructor() {
        owner = msg.sender;
    }

    // modifier which restricts function to owner of contract
    modifier onlyOwner() {
        require(msg.sender == owner, "Only the owner can call this function.");
        _;
    }

    // modifier which restricts function to owner of contract and authorized addresses
    modifier onlyAuthorized() {
        require(authorizedAddresses[msg.sender] || msg.sender == owner, "Only authorized addresses can call this function.");
        _;
    }

    // gives access to change temperature to new address
    function addAuthorizedAddress(address _address) public onlyOwner {
        authorizedAddresses[_address] = true;
    }

    // removes access to change temperature from address
    function removeAuthorizedAddress(address _address) public onlyOwner {
        authorizedAddresses[_address] = false;
    }

    // sets temperature to value
    function setTemperature(uint _temperature) public onlyAuthorized {
        temperature = _temperature;
    }
}