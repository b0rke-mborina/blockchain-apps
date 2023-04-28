pragma solidity >=0.7.0 <0.9.0;

// T-Mobile razvija ovaj ugovor
// Pratit subscription samo za jednog korisnika


contract CellSubscription {
    uint monthlyCost;
    address private owner;

    // pri pokretanju ugovora, samo jednom
    constructor(uint cost) {
        monthlyCost = cost;
        owner = this.address;
    }

    function makePayment() public payable {}

    receive() external payable {}
    
    function isPaid() public view returns (bool) {
        return monthlyCost <= address(this).balance;
    }

    function withdraw() public {
        // prenesi novac s ugovora na onog tko je napravio zahtjev za withdraw funkciju
        // Ako sam T-Mobile tek onda transfer, inače ništa
        require(msg.sender == owner, "")
        if (owner == msg.sender)
            payable(msg.sender).transfer(address(this).balance);
    }
}
