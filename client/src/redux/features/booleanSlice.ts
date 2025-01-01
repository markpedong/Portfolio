import { createSlice } from '@reduxjs/toolkit';

type InitialSlice = {
	darkMode: boolean;
};
const initialState: InitialSlice = {
	darkMode: false
};

export const BooleanSlice = createSlice({
	name: 'Boolean',
	initialState,
	reducers: {
		setDarkMode: state => ({ ...state, darkMode: !state.darkMode })
	}
});

export const { setDarkMode } = BooleanSlice.actions;
export default BooleanSlice.reducer;
