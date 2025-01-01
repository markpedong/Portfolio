import { ModalFormProps, ProFormProps, ProTableProps } from '@ant-design/pro-components'
import { Rule } from 'antd/es/form'

export const INPUT_LINK: Rule = {
	validator: (_, value) => {
		if (!value || /^(https?|ftp):\/\/[^\s/$.?#].[^\s]*$/.test(value)) {
			return Promise.resolve()
		}
		return Promise.reject('Please enter a valid link starting with http:// or https://')
	}
}

export const INPUT_EMAIL: Rule = {
	type: 'email',
	message: 'Please enter a valid Email'
}

export const PRO_TABLE_PROPS: ProTableProps<any, any> = {
	options: false,
	scroll: { x: 1200 },
	search: {
		labelWidth: 'auto',
		collapsed: false,
		collapseRender: false,
		span: {
			xs: 24,
			sm: 12,
			md: 8,
			lg: 6,
			xl: 6,
			xxl: 4
		}
	},
	pagination: {
		defaultPageSize: 20,
		showSizeChanger: true
	}
}

export const MODAL_FORM_PROPS: ModalFormProps = {
	labelCol: { flex: '110px' },
	layout: 'horizontal',
	width: 800,
	modalProps: {
		destroyOnClose: true,
		maskClosable: false
	},
	wrapperCol: { flex: 1 },
	autoFocusFirstInput: true,
	preserve: false,
	size: 'large'
}

export const FORM_PROPS: ProFormProps = {
	labelCol: { flex: '80px' },
	layout: 'horizontal',
	wrapperCol: { flex: 1 },
	preserve: false,
	size: 'middle'
}
