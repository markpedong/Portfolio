import { createSlice } from '@reduxjs/toolkit'

type InitialSlice = {
	userData: {
		id: string
		created_at: string
		updated_at: string
		password: string
		email: string
		first_name: string
		birthday: string
		username: string
	}
	type: number
	token: string
	isLoggedIn: boolean
}
const initialState: InitialSlice = {
	userData: {
		id: '',
		created_at: '',
		updated_at: '',
		password: '',
		email: '',
		first_name: '',
		birthday: '',
		username: ''
	},
	type: 0,
	token: '',
	isLoggedIn: false
}

export const UserSlice = createSlice({
	name: 'Boolean',
	initialState,
	reducers: {
		resetUserData: () => initialState,
		setType: (state, action) => ({ ...state, type: action.payload }),
		setUserData: (state, action) => ({ ...state, isLoggedIn: true, userData: action.payload }),
		setUserToken: (state, action) => ({ ...state, token: action.payload })
	}
})

export const { setUserData, resetUserData, setUserToken, setType } = UserSlice.actions
export default UserSlice.reducer
