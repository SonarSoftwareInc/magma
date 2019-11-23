/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import RelayEnvironment from '../common/RelayEnvironment.js';
import {commitMutation, graphql} from 'react-relay';
import type {
  ExecuteWorkOrderMutationResponse,
  ExecuteWorkOrderMutationVariables,
} from './__generated__/ExecuteWorkOrderMutation.graphql';
import type {MutationCallbacks} from './MutationCallbacks.js';

const mutation = graphql`
  mutation ExecuteWorkOrderMutation($id: ID!) {
    executeWorkOrder(id: $id) {
      equipmentAdded {
        ...EquipmentTable_equipment
      }
      equipmentRemoved
      linkAdded {
        ...EquipmentPortsTable_link
      }
      linkRemoved
    }
  }
`;

export default (
  variables: ExecuteWorkOrderMutationVariables,
  callbacks?: MutationCallbacks<ExecuteWorkOrderMutationResponse>,
  updater?: (store: any) => void,
) => {
  const {onCompleted, onError} = callbacks ? callbacks : {};
  commitMutation(RelayEnvironment, {
    mutation,
    variables,
    updater,
    onCompleted,
    onError,
  });
};