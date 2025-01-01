import { useAppSelector } from '@/redux/store'
import { CaretDownOutlined, CaretUpOutlined } from '@ant-design/icons'
import { ConfigProvider, message, theme } from 'antd'
import React, { MutableRefObject } from 'react'
import enUS from 'antd/locale/en_US'
import { ActionType, ProFormInstance } from '@ant-design/pro-components'
import { ApiResponse } from '@/api/http'
import { FileType } from '@/api'

export const afterModalformFinish = (
	actionRef?: MutableRefObject<ActionType | undefined>,
	res?: ApiResponse<any>,
	formRef?: MutableRefObject<ProFormInstance | undefined>
) => {
	if (actionRef) {
		actionRef?.current?.reload()
	}

	if (formRef) {
		formRef?.current?.setFieldsValue(res?.data.data)
	}

	if (res?.data?.success) {
		message.success(res?.data?.message)
	}

	return !!res?.data?.success
}

export const renderPercentage = percentage => {
	const per = percentage?.toFixed(2)

	return (
		<span
			style={{
				color: per > 0.0 ? '#16c784' : '#ea3943'
			}}
		>
			{per > 0.0 ? <CaretUpOutlined /> : per < 0.0 ? <CaretDownOutlined /> : !per && ''} {per?.replace('-', '')}
		</span>
	)
}

export const AntdConfigProvider = ({ children }: { children: React.ReactNode }) => {
	const { darkMode } = useAppSelector(s => s.boolean)

	return (
		<ConfigProvider
			locale={enUS}
			theme={{
				algorithm: darkMode ? theme.darkAlgorithm : theme.defaultAlgorithm,
				token: { colorPrimary: 'red', fontFamily: 'Poppins' }
			}}
		>
			{children}
		</ConfigProvider>
	)
}

export const BeforeUpload = (file: FileType) => {
	const isJpgOrPngOrSvg = ["image/jpeg", "image/png", "image/svg+xml", "image/webp"].includes(file.type)
	if (!isJpgOrPngOrSvg) {
		message.error('You can only upload JPG/PNG/SVG/WEBP file!')
		return
	}
	const isLt2M = file.size / 1024 / 1024 < 2
	if (!isLt2M) {
		message.error('Image must smaller than 2MB!')
		return
	}
	return isJpgOrPngOrSvg && isLt2M
}
