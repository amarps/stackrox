import React from 'react';
import { connect } from 'react-redux';
import { createStructuredSelector } from 'reselect';
import * as Icon from 'react-feather';

import { selectors } from 'reducers';
import { actions as pageActions } from 'reducers/network/page';
import { timeWindows } from 'constants/timeWindows';

type TimeWindowSelectorProps = {
    setActivityTimeWindow: (timeWindow) => void;
    activityTimeWindow: string;
    isDisabled?: boolean;
};

function TimeWindowSelector({
    setActivityTimeWindow,
    activityTimeWindow,
    isDisabled = false,
}: TimeWindowSelectorProps) {
    function selectTimeWindow(event) {
        const timeWindow = event.target.value;
        setActivityTimeWindow(timeWindow);
    }

    return (
        <div
            className={`flex relative whitespace-nowrap border-2 rounded-sm mr-2 ml-2 min-h-10 bg-base-100 border-base-300 hover:border-base-400 ${
                isDisabled ? 'disabled' : ''
            }`}
        >
            <div className="absolute inset-y-0 ml-2 flex items-center cursor-pointer z-0 pointer-events-none">
                <Icon.Clock className="h-4 w-4 text-base-500" />
            </div>
            <select
                className="pl-8 pr-8 truncate text-lg bg-base-100 py-2 text-sm text-base-600 border-0 hover:border-base-300 cursor-pointer"
                onChange={selectTimeWindow}
                value={activityTimeWindow}
                disabled={isDisabled}
            >
                {timeWindows.map((window) => (
                    <option key={window} value={window}>
                        {window}
                    </option>
                ))}
            </select>
        </div>
    );
}

const mapStateToProps = createStructuredSelector({
    activityTimeWindow: selectors.getNetworkActivityTimeWindow,
});

const mapDispatchToProps = {
    setActivityTimeWindow: pageActions.setNetworkActivityTimeWindow,
};

export default connect(mapStateToProps, mapDispatchToProps)(TimeWindowSelector);