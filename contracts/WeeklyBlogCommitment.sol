// SPDX-License-Identifier: MIT
pragma solidity >=0.8.4;

import "./AccessControl.sol";
import "./RssFeedLengthOracle.sol";

contract WeeklyBlogCommitment is AccessControl {
    struct Commitment {
        address payable fulfillAddress;
        address payable breachAddress;
        string rssFeed;
        uint balance;
        uint last_checked;
        uint8 factor;
    }

    event NewCommitment(string rss_feed);

    Commitment[] private commitments;

    mapping(address => uint) private pendingWithdrawals;

    RssFeedLengthOracle oracleContract;

    function setOracleAddress(address addr) external onlyAdmin {
        oracleContract = RssFeedLengthOracle(addr);
    }

    function getCommitments() external view onlyAdmin returns (Commitment[] memory) {
        return commitments;
    }

    function commit(string memory rssFeed, uint8 factor, address payable breachAddress) external payable {
        require(msg.value > 0);
        require(factor > 0);
        commitments.push(
            Commitment(payable(msg.sender), breachAddress, rssFeed, msg.value, 0, factor)
        );
        emit NewCommitment(rssFeed);
    }

    function resolve() external onlyAdmin {
        uint _new_commitments_size = 0;

        // First calculate the number of commitments that will be checked
        for (uint i = 0; i < commitments.length; i = unchecked_inc(i)) {
            Commitment memory commitment = commitments[i];

            if (shouldCheckCommitment(commitment)) {
                unchecked { _new_commitments_size += 1; }
            }
        }
        
        address payable[] memory pending_addrs = new address payable[](_new_commitments_size);
        uint[] memory pending_amounts = new uint[](_new_commitments_size);

        for (uint i = 0; i < commitments.length; i = unchecked_inc(i)) {
            Commitment storage commitment = commitments[i];

            if (!shouldCheckCommitment(commitment)) {
                continue;
            }

            unchecked {
               commitment.last_checked = block.timestamp;

               uint commitment_cost = commitment.balance / commitment.factor;
               commitment_cost = commitment_cost > commitment.balance ? commitment.balance : commitment_cost;

               // There will be a reamining balance after this.
               if (commitment.balance > commitment_cost) {
                   commitment.balance -= commitment_cost;
               }

               address payable to = wasFulfilled(commitment) ? 
                   commitment.fulfillAddress :
                   commitment.breachAddress;

               pending_addrs[i] = to;
               pending_amounts[i] = commitment_cost;
            }
        }

        for (uint i = 0; i < _new_commitments_size; unchecked_inc(i)) {
            address payable to = pending_addrs[i];
            uint amount = pending_amounts[i];

            if (amount == 0) break;

            // Send commitment balance to appropriate address and make available for 
            // withdrawl if failed.
            if (!to.send(amount)) {
                unchecked {
                    uint current_amount = pendingWithdrawals[to];
                    uint new_amount = current_amount + amount;
                    pendingWithdrawals[to] = new_amount > current_amount ? new_amount : (2**256 - 1);
                }
            }
        }
    }

    function shouldCheckCommitment(Commitment memory commitment) private view returns (bool) {
        bool last_checked_within_a_week = (block.timestamp - (1 weeks)) <= commitment.last_checked;
        return commitment.last_checked > 0 && commitment.balance > 0 && !last_checked_within_a_week;
    }

    function withdraw() external returns (bool) { 
        uint amount = pendingWithdrawals[msg.sender];
        if (amount > 0) {
            pendingWithdrawals[msg.sender] = 0;

            if (!payable(msg.sender).send(amount)) {
                pendingWithdrawals[msg.sender] = amount;
                return false;
            }
        }
        return true;
    }

    function amount_available_to_withdraw() external view returns (uint) {
        return pendingWithdrawals[msg.sender];
    }

    function wasFulfilled(Commitment memory commitment) private view returns(bool) {
        uint change = oracleContract.getLengthChangeWitinRange(
            commitment.rssFeed,
            block.timestamp - (1 weeks),
            block.timestamp);

        return change > 0;
    }

    function unchecked_inc(uint i) private pure returns (uint) {
        unchecked { return i + 1; }
    }
}