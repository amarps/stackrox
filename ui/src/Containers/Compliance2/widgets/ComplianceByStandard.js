import React from 'react';
import PropTypes from 'prop-types';
import componentTypes from 'constants/componentTypes';
import standardLabels from 'messages/standards';

import Widget from 'Components/Widget';
import Sunburst from 'Components/visuals/Sunburst';
import Query from 'Components/AppQuery';
import Loader from 'Components/Loader';

const MAX_CHAR = 100;

const colors = [
    'var(--tertiary-400)',
    'var(--warning-400)',
    'var(--caution-400)',
    'var(--alert-400)'
];
const getColor = value => {
    if (value === 100) return colors[0];
    if (value >= 70) return colors[1];
    if (value >= 50) return colors[2];
    return colors[3];
};

const sunburstLegendData = [
    { title: '100%', color: 'var(--tertiary-400)' },
    { title: '> 70%', color: 'var(--warning-400)' },
    { title: '> 50%', color: 'var(--caution-400)' },
    { title: '< 50%', color: 'var(--alert-400)' }
];

const processSunburstData = (data, type) => {
    const groupStatsMapping = {};
    const controlStatsMapping = {};
    const groupMapping = {};

    data.groupResults.results.forEach(({ aggregationKeys, numPassing, numFailing }) => {
        groupStatsMapping[`${aggregationKeys[0].id}`] =
            Math.round((numPassing / (numFailing + numPassing)) * 100) || 0;
    });

    data.controlResults.results.forEach(({ aggregationKeys, numPassing, numFailing }) => {
        controlStatsMapping[`${aggregationKeys[0].id}`] =
            Math.round((numPassing / (numFailing + numPassing)) * 100) || 0;
    });

    data.complianceStandards
        .filter(datum => datum.id.includes(type))
        .forEach(({ groups, controls }) => {
            groups
                .filter(group => group.standardId.includes(type))
                .forEach(datum => {
                    const value = groupStatsMapping[datum.id] || 0;
                    groupMapping[datum.id] = {
                        name: `${datum.name}. ${datum.description.substring(0, MAX_CHAR)}${
                            datum.description.length > MAX_CHAR ? '...' : ''
                        }`,
                        color: getColor(value),
                        value,
                        children: []
                    };
                });
            controls
                .filter(control => control.standardId.includes(type))
                .forEach(datum => {
                    if (groupMapping[datum.groupId]) {
                        const group = groupMapping[datum.groupId];
                        const value = controlStatsMapping[`${datum.standardId}:${datum.id}`] || 0;
                        group.children.push({
                            name: `${datum.name} - ${datum.description.substring(0, MAX_CHAR)}${
                                datum.description.length > MAX_CHAR ? '...' : ''
                            }`,
                            color: getColor(value),
                            link: `main/compliance2/${datum.standardId}/${datum.id}`,
                            value
                        });
                    }
                });
        });
    return Object.values(groupMapping);
};

const getNumControls = sunburstData =>
    sunburstData.reduce((acc, curr) => acc + curr.children.length, 0);

const ComplianceByStandard = ({ type, params }) => (
    <Query params={params} componentType={componentTypes.COMPLIANCE_BY_STANDARD}>
        {({ loading, data }) => {
            let contents = <Loader />;
            const headerText = `${standardLabels[type]} Compliance`;
            if (!loading && data) {
                const sunburstData = processSunburstData(data, type);
                const sunburstRootData = [
                    {
                        text: `${sunburstData.length} Categories`
                    },
                    {
                        text: `${getNumControls(sunburstData)} Controls`
                    }
                ];

                if (!sunburstData.length) {
                    contents = (
                        <>
                            <div className="flex flex-1 items-center justify-center">
                                No Data Available
                            </div>
                        </>
                    );
                } else {
                    contents = (
                        <Sunburst
                            data={sunburstData}
                            rootData={sunburstRootData}
                            legendData={sunburstLegendData}
                        />
                    );
                }
            }

            return <Widget header={headerText}>{contents}</Widget>;
        }}
    </Query>
);

ComplianceByStandard.propTypes = {
    type: PropTypes.string.isRequired,
    params: PropTypes.shape({})
};

ComplianceByStandard.defaultProps = {
    params: null
};

export default ComplianceByStandard;
