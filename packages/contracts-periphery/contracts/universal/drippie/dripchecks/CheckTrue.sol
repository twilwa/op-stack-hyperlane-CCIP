// SPDX-License-Identifier: MIT
pragma solidity 0.8.16;

import { IDripCheck } from "../IDripCheck.sol";

/**
 * @title CheckTrue
 * @notice DripCheck that always returns true.
 */
contract CheckTrue is IDripCheck {
    /**
     * @inheritdoc IDripCheck
     */
    function check(bytes memory) external pure returns (bool) {
        return true;
    }
}
