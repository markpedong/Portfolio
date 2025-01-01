import { resetUserData } from '@/redux/features/userSlice'
import { useAppDispatch } from '@/redux/store'
import { getLocalStorage } from '@/utils/xLocalstorage'
import { useEffect } from 'react'
import { useNavigate } from 'react-router-dom'

export default function isAuth(Component: any) {
	const dispatch = useAppDispatch()
	const navigate = useNavigate()
	const auth = getLocalStorage('refresh_token')

	useEffect(() => {
		if (!auth) {
			dispatch(resetUserData())
			navigate('/')
		}
	}, [auth, dispatch, navigate])

	return Component
}
