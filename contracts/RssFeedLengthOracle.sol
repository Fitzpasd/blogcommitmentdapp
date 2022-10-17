// SPDX-License-Identifier: MIT
pragma solidity >=0.8.12;

import "./AccessControl.sol";

contract RssFeedLengthOracle is AccessControl {
    struct FeedInfo { 
        uint date;
        uint length;
    }

    mapping(string => FeedInfo[]) feedInfoMap;

    error FeedNotTracked(string url);
    error NotEnoughDataSaved();
    error DatesNotTracked();
    error UrlIsEmpty();

    modifier feedIsTracked(string memory url) {
        FeedInfo[] memory info = feedInfoMap[url];

        if (info.length == 0)
            revert FeedNotTracked(url);

        _;
    }

    modifier nonEmptyUrl(string memory url) {
        if(bytes(url).length == 0)
            revert UrlIsEmpty();

        _;
    }

    function getAllEntriesForUrl(string memory url) external onlyAdmin nonEmptyUrl(url) feedIsTracked(url) view returns (FeedInfo[] memory) {
        return feedInfoMap[url];
    }

    function getLengthOfFeed(string memory url) external view nonEmptyUrl(url) feedIsTracked(url) returns (uint) {
        FeedInfo[] memory info = feedInfoMap[url];
        return feedInfoMap[url][info.length - 1].length;
    }

    function getLengthChangeWitinRange(string memory url, uint start, uint end) external view nonEmptyUrl(url) feedIsTracked(url) returns (uint) {
        unchecked {
            require(start < end);

            FeedInfo[] memory info = feedInfoMap[url];

            if (info.length < 2)
                revert NotEnoughDataSaved();

            if (info[0].date > end || info[info.length - 1].date < start)
                revert DatesNotTracked();

            FeedInfo memory startInfo;
            FeedInfo memory endInfo;

            for (uint i = 0; i < info.length; i++) {
                if (info[i].date >= start) {
                    startInfo = info[i];
                    break;
                }
            }

            for (uint i = info.length; i > 0; i--) {
                if (info[i-1].date <= end) {
                    endInfo = info[i-1];
                    break;
                }
            }

            return endInfo.length - startInfo.length;
        }
    }

    function setCurrentLengthOfFeed(string memory url, uint length) external onlyAdmin nonEmptyUrl(url) {
        require(length >= 0);
        feedInfoMap[url].push(FeedInfo(block.timestamp, length));
    }
}