import Root from '@/pages/root'
import '@/styles/global.scss'
import ReactDOM from 'react-dom/client'
import { BrowserRouter } from 'react-router-dom'
import ReduxProvider from './redux/provider'
import { AntdConfigProvider } from './utils/antd'
import { ignoreFindDOMNodeError } from './utils'

ignoreFindDOMNodeError()

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <BrowserRouter>
    <ReduxProvider>
      <AntdConfigProvider>
        <Root />
      </AntdConfigProvider>
    </ReduxProvider>
  </BrowserRouter>
)
