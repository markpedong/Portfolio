import React from 'react'
import { Provider } from 'react-redux'
import { persistStore } from 'redux-persist'
import { PersistGate } from 'redux-persist/integration/react'
import { store } from './store'
let persistor = persistStore(store)

const ReduxProvider = ({ children }: { children: React.ReactNode }) => {
	return (
		<Provider store={store}>
			<PersistGate persistor={persistor}>{children}</PersistGate>
		</Provider>
	)
}

export default ReduxProvider
