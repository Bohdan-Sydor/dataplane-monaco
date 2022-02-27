import { Grid } from '@mui/material';
import CustomChip from '../../components/CustomChip';
import { useGlobalFlowState } from '../Flow';
import { useGlobalRunState } from './useWebSocket';

export default function StatusChips() {
    // Global state
    const RunState = useGlobalRunState();
    const FlowState = useGlobalFlowState();

    return (
        <Grid item display="flex" alignItems="center" sx={{ alignSelf: 'center' }} flex={1.2}>
            <CustomChip amount={Object.values(FlowState.elements.get()).filter((a) => a.type !== 'custom').length} label="Steps" margin={2} customColor="orange" />
            <CustomChip amount={Object.values(RunState.get()).filter((a) => a === 'Queue').length} label="Queue" margin={2} customColor="purple" />
            <CustomChip amount={Object.values(RunState.get()).filter((a) => a === 'Run').length} label="Running" margin={2} customColor="blue" />
            <CustomChip amount={Object.values(RunState.get()).filter((a) => a === 'Success').length} label="Succeeded" margin={2} customColor="green" />
            <CustomChip amount={Object.values(RunState.get()).filter((a) => a === 'Fail').length} label="Failed" margin={2} customColor="red" />
        </Grid>
    );
}
