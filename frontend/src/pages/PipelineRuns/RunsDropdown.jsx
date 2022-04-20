import { Autocomplete, Grid, TextField } from '@mui/material';
import React, { useEffect, useState } from 'react';
import { useGetPipelineRuns } from '../../graphql/getPipelineRuns';
import { useHistory, useParams } from 'react-router-dom';
import { useSnackbar } from 'notistack';
import { usePipelineTasksRun } from '../../graphql/getPipelineTasksRun';
import { useGlobalPipelineRun} from './GlobalPipelineRunUIState'
import { useGlobalRunState } from './GlobalRunState';
import { useGetPipelineFlow } from '../../graphql/getPipelineFlow';
import { prepareInputForFrontend } from '.';
import useOnChangeDropdownWebSocket from './useOnChangeDropdownWebSocket';
import useOnPageLoadWebSocket from './useOnPageLoadWebSocket';
import { formatDateNoZone } from '../../utils/formatDate';
import { useGlobalMeState } from '../../components/Navbar';


/*
Input:
runs, setRuns - state set at RunNavBar level but are set inside this component
*/
export default function RunsDropdown({ environmentID, pipeline, runs, setRuns, selectedRun, setSelectedRun }) {
    // Global states
    const RunState = useGlobalPipelineRun();
    const MeData = useGlobalMeState();

    // Local state
    const [isNewFlow, setIsNewFlow] = useState(false);

    // Graphql
    const getPipelineRuns = useGetPipelineRuns();
    // const getPipelineRuns = useGetPipelineRunsHook(environmentID, setRuns, setIsNewFlow, pipeline.updated_at);
    // const getPipelineFlow = useGetPipelineFlowHook(pipeline.pipelineID);

    const { enqueueSnackbar } = useSnackbar();

    // Instantiate websocket for on page load
    // useOnPageLoadWebSocket(environmentID, setSelectedRun, setRuns, setIsNewFlow, pipeline.updated_at);

    // Instantiate websocket for dropdown change
    // useOnChangeDropdownWebSocket(environmentID, setSelectedRun);

    useEffect(() => {
        (async () => {

            // --------  1. Retrieve the the previous runs:
            const response = await getPipelineRuns({ pipelineID: pipeline.pipelineID, environmentID });

            console.log("Runs response:", response)
    
            if (response.length === 0) {
                setRuns([]);
                // If there are no runs to get flow info, run getPipelineFlow instead
                // getPipelineFlow(environmentID);
            } else if (response.r || response.error) {
                enqueueSnackbar("Can't get runs: " + (response.msg || response.r || response.error), { variant: 'error' });
            } else if (response.errors) {
                response.errors.map((err) => enqueueSnackbar(err.message, { variant: 'error' }));
            } else {
                setRuns(response);
                // return [response[0]?.updated_at, response[0].run_id];
            }

            let lastRunTime = response[0]?.updated_at
            let runID = response[0].run_id

            // On page load select the latest response
            setSelectedRun(response[0])

            // Get the flow of the latest run or if no flow then get structure



            // 2. Retrieve the latest run

            // If the pipeline has a new flow, get only the flow and return
            const isNewFlow = pipeline.updated_at > lastRunTime;
            if (isNewFlow) {
                setIsNewFlow(true);
                // getPipelineFlow(environmentID);
                return;
            }

            RunState.merge({
                selectedRunID: runID,
                onLoadTrigger: 1,
            });
        })();

        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, []);

    // Update elements on run dropdown change
    const handleDropdownChange = (run) => {
        setSelectedRun(run);
        RunState.merge((r) => ({
            selectedRunID: run.run_id,
            onChangeTrigger: r.onChangeTrigger + 1,
        }));
    };

    return (
        <Grid item alignItems="center" display="flex" width={520}>
            {selectedRun && !isNewFlow ? (
                <Autocomplete
                    id="run_autocomplete"
                    onChange={(event, newValue) => handleDropdownChange(newValue)}
                    value={selectedRun}
                    disableClearable
                    sx={{ minWidth: '520px' }}
                    options={runs}
                    isOptionEqualToValue={(option, value) => option.run_id === value.run_id}
                    getOptionLabel={(a) => formatDateNoZone(a.created_at, MeData.timezone.get()) + ' - ' + a.run_id}
                    renderInput={(params) => <TextField {...params} label="Run" id="run" size="small" sx={{ fontSize: '.75rem', display: 'flex' }} />}
                />
            ) : null}
            {isNewFlow ? (
                <Autocomplete
                    id="run_autocomplete"
                    onChange={(event, newValue) => handleDropdownChange(newValue)}
                    value={selectedRun}
                    disableClearable
                    sx={{ minWidth: '520px' }}
                    options={runs}
                    isOptionEqualToValue={(option, value) => option.run_id === value.run_id}
                    getOptionLabel={(a) => formatDateNoZone(a.created_at, MeData.timezone.get()) + ' - ' + a.run_id}
                    renderInput={(params) => <TextField {...params} label="Run" id="run" size="small" sx={{ fontSize: '.75rem', display: 'flex' }} />}
                />
            ) : null}
        </Grid>
    );
}




// const useGetPipelineFlowHook = () => {
//     // GraphQL hook
//     const getPipelineFlow = useGetPipelineFlow();

//     // React router
//     const history = useHistory();

//     // Global state
//     const FlowState = useGlobalPipelineRun();

//     // URI parameter
//     const { pipelineId } = useParams();

//     const { enqueueSnackbar, closeSnackbar } = useSnackbar();

//     // Get members
//     return async (environmentID) => {
//         const rawResponse = await getPipelineFlow({ pipelineID: pipelineId, environmentID });
//         const response = prepareInputForFrontend(rawResponse);

//         if (response.length === 0) {
//             FlowState.elements.set([]);
//             history.push(`/pipelines/flow/${pipelineId}`);
//         } else if (response.r === 'error') {
//             closeSnackbar();
//             enqueueSnackbar("Can't get flow: " + response.msg, { variant: 'error' });
//         } else if (response.errors) {
//             response.errors.map((err) => enqueueSnackbar(err.message, { variant: 'error' }));
//         } else {
//             FlowState.elements.set(response);
//         }
//     };
// };


