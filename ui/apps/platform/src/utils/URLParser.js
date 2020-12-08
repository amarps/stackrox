import { matchPath } from 'react-router-dom';
import qs from 'qs';

import useCases from 'constants/useCaseTypes';
import { searchParams, sortParams, pagingParams } from 'constants/searchParams';
import WorkflowEntity from './WorkflowEntity';
import { WorkflowState } from './WorkflowState';
import {
    nestedPaths as workflowPaths,
    urlEntityListTypes,
    urlEntityTypes,
    clustersPathWithParam,
    riskPath,
} from '../routePaths';

function getTypeKeyFromParamValue(value, listOnly) {
    const listMatch = Object.entries(urlEntityListTypes).find((entry) => entry[1] === value);
    const entityMatch = Object.entries(urlEntityTypes).find((entry) => entry[1] === value);
    const match = listOnly ? listMatch : listMatch || entityMatch;
    return match ? match[0] : null;
}

function getEntityFromURLParam(type, id) {
    return new WorkflowEntity(getTypeKeyFromParamValue(type), id);
}

function paramsToStateStack(params) {
    const { pageEntityListType, pageEntityType, pageEntityId, entityId1, entityId2 } = params;
    const { entityType1: urlEntityType1, entityType2: urlEntityType2 } = params;
    const entityListType1 = getTypeKeyFromParamValue(urlEntityType1, true);
    const entityListType2 = getTypeKeyFromParamValue(urlEntityType2, true);
    const entityType1 = getTypeKeyFromParamValue(urlEntityType1);
    const entityType2 = getTypeKeyFromParamValue(urlEntityType2);
    const stateArray = [];
    if (!pageEntityListType && !pageEntityType) {
        return stateArray;
    }

    // List
    if (pageEntityListType) {
        stateArray.push(getEntityFromURLParam(pageEntityListType));

        if (entityId1) {
            stateArray.push(getEntityFromURLParam(pageEntityListType, entityId1));
        }
    } else {
        stateArray.push(getEntityFromURLParam(pageEntityType, pageEntityId));
        if (entityListType1) {
            stateArray.push(new WorkflowEntity(entityListType1));
        }
        if (entityType1 && entityId1) {
            stateArray.push(new WorkflowEntity(entityType1, entityId1));
        }
    }

    if (entityListType2) {
        stateArray.push(new WorkflowEntity(entityListType2));
    }
    if (entityType2 && entityId2) {
        stateArray.push(new WorkflowEntity(entityType2, entityId2));
    }

    return stateArray;
}

function formatSort(sort) {
    if (!sort) {
        return null;
    }

    let sorts;
    if (!Array.isArray(sort)) {
        sorts = [sort];
    } else {
        sorts = [...sort];
    }

    return sorts.map(({ id, desc }) => {
        return {
            id,
            desc: JSON.parse(desc),
        };
    });
}

// Convert URL to workflow state and search objects
// note: this will read strictly from 'location' as 'match' is relative to the closest Route component
function parseURL(location) {
    if (!location) {
        return {};
    }

    const { pathname, search } = location;
    const listParams = matchPath(pathname, {
        path: workflowPaths.LIST,
    });
    const entityParams = matchPath(pathname, {
        path: workflowPaths.ENTITY,
    });
    const dashboardParams = matchPath(pathname, {
        path: workflowPaths.DASHBOARD,
        exact: true,
    });

    // check for legacy-Workflow sections
    const matchedRiskParams = matchPath(pathname, {
        path: riskPath,
        exact: true,
    });
    const matchedClustersParams = matchPath(pathname, {
        path: clustersPathWithParam,
        exact: true,
    });
    const legacyParams = matchedRiskParams
        ? {
              params: {
                  ...matchedRiskParams,
                  context: useCases.RISK,
              },
          }
        : {
              params: {
                  ...matchedClustersParams,
                  context: useCases.CLUSTERS,
              },
          };

    const { params } = entityParams || listParams || dashboardParams || legacyParams;
    const queryStr = search ? qs.parse(search, { ignoreQueryPrefix: true }) : {};

    const stateStackFromURLParams = paramsToStateStack(params) || [];

    let { workflowState: stateStackFromQueryString = [] } = queryStr;
    const {
        [searchParams.page]: pageSearch,
        [searchParams.sidePanel]: sidePanelSearch,
        [sortParams.page]: pageSort,
        [sortParams.sidePanel]: sidePanelSort,
        [pagingParams.page]: pagePaging,
        [pagingParams.sidePanel]: sidePanelPaging,
    } = queryStr;

    stateStackFromQueryString = !Array.isArray(stateStackFromQueryString)
        ? [stateStackFromQueryString]
        : stateStackFromQueryString;
    stateStackFromQueryString = stateStackFromQueryString.map(
        ({ t, i }) => new WorkflowEntity(t, i)
    );

    const workflowState = new WorkflowState(
        params.context,
        [...stateStackFromURLParams, ...stateStackFromQueryString],
        {
            [searchParams.page]: pageSearch || null,
            [searchParams.sidePanel]: sidePanelSearch || null,
        },
        {
            [sortParams.page]: formatSort(pageSort),
            [sortParams.sidePanel]: formatSort(sidePanelSort),
        },
        {
            [pagingParams.page]: parseInt(pagePaging || 0, 10),
            [pagingParams.sidePanel]: parseInt(sidePanelPaging || 0, 10),
        }
    );

    return workflowState;
}

export default parseURL;
