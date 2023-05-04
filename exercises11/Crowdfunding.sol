// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0 <0.9.0;

contract Crowdfunding {
    // campaign structure
    struct Campaign {
        address payable creator;
        string title;
        uint256 goal;
        uint256 raisedAmount;
        uint256 deadline;
    }
    
    // array of campaigns
    Campaign[] public campaigns;
    
    // creates a new campaign with the given parameters and adds it to an array of campaigns.
    function createCampaign(string memory _title, uint256 _goal, uint256 _deadline) public {
        Campaign memory newCampaign = Campaign({
            creator: payable(msg.sender),
            title: _title,
            goal: _goal,
            raisedAmount: 0,
            deadline: block.timestamp + _deadline
        });
        campaigns.push(newCampaign);
    }
    
    // allow users to contribute to a campaign by sending ether to the contrac
    // the amount contributed should be added to the raisedAmount field of the Campaign struct for the given campaign index
    function contribute(uint256 _campaignIndex) public payable campaignExists(_campaignIndex) campaignDeadlineNotPassed(_campaignIndex) {
        Campaign storage campaign = campaigns[_campaignIndex];
        campaign.raisedAmount += msg.value;
    }
    
    // allows the creator of a campaign to withdraw the raised funds
    // if the campaign has met its funding goal and the deadline has passed
    function withdrawFunds(uint256 _campaignIndex) public onlyCreator(_campaignIndex) campaignExists(_campaignIndex) campaignGoalReached(_campaignIndex) {
        Campaign storage campaign = campaigns[_campaignIndex];
        campaign.creator.transfer(campaign.raisedAmount);
        campaign.raisedAmount = 0;
    }
    
    // ensures that only the creator of a campaign can call certain functions
    modifier onlyCreator(uint256 _campaignIndex) {
        require(campaigns[_campaignIndex].creator == msg.sender, "Only the campaign creator can perform this action");
        _;
    }
    
    // ensures that the campaign index being passed to a function corresponds to an existing campaign
    modifier campaignExists(uint256 _campaignIndex) {
        require(_campaignIndex < campaigns.length, "Campaign does not exist");
        _;
    }
    
    // ensures that the campaign funding goal is reached
    modifier campaignGoalReached(uint256 _campaignIndex) {
        Campaign storage campaign = campaigns[_campaignIndex];
        require(campaign.raisedAmount >= campaign.goal, "Campaign goal not reached");
        _;
    }
    
    // ensures that the campaign deadline has passed
    modifier campaignDeadlineNotPassed(uint256 _campaignIndex) {
        Campaign storage campaign = campaigns[_campaignIndex];
        require(block.timestamp <= campaign.deadline, "Campaign deadline has passed");
        _;
    }
}