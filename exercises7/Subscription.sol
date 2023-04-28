pragma solidity >=0.7.0 <0.9.0;

// T-Mobile razvija ovaj ugovor
// Pratit subscription samo za jednog korisnika


contract Subscription {
    mapping(address => uint) public balances;

    function updateBalance(uint newBalance) public {
        balances[msg.sender] = newBalance;
    }

    // pri pokretanju ugovora, samo jednom
    constructor(uint cost) {
        
    }

    function makePayment() public payable {}

    receive() external payable {}
}
