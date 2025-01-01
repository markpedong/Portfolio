import {
	addApplication,
	deleteApplication,
	getApplication,
	TApplicationItem,
	toggleAppStatus,
	updateApplication,
	uploadImage
} from '@/api'
import { GLOBAL_STATUS } from '@/api/constants'
import { MODAL_FORM_PROPS, PRO_TABLE_PROPS } from '@/constants'
import { dateTimeFormatter } from '@/utils'
import { afterModalformFinish, BeforeUpload } from '@/utils/antd'
import {
	ActionType,
	ModalForm,
	ProColumns,
	ProFormText,
	ProFormUploadButton,
	ProTable
} from '@ant-design/pro-components'
import { Button, Image, Popconfirm, Space, Switch, Typography } from 'antd'
import { omit } from 'lodash'
import { useRef, useState } from 'react'

const Application = () => {
	const [imgUrl, setImgUrl] = useState('')
	const actionRef = useRef<ActionType>()
	const columns: ProColumns<TApplicationItem>[] = [
		{
			title: 'Name',
			align: 'center',
			dataIndex: 'name'
		},
		{
			title: 'Logo',
			align: 'center',
			search: false,
			render: (_, record) => <Image src={record?.image} alt="image" width={60} height={60} />
		},
		{
			title: (
				<div className="flex flex-col gap-0">
					<div>Created</div>
					<div>Updated</div>
				</div>
			),
			search: false,
			align: 'center',
			render: (_, record) => (
				<div className="flex flex-col">
					<div>{dateTimeFormatter(record.created_at, 'MM-DD-YYYY HH:MM:ss')}</div>
					<div>{dateTimeFormatter(record.updated_at, 'MM-DD-YYYY HH:MM:ss')}</div>
				</div>
			)
		},
		{
			title: 'Operator',
			align: 'center',
			search: false,
			render: (_, record) => (
				<Space>
					{renderSwitch(record)}
					{renderAddEditApp('EDIT', record)}
					{renderDeleteApp(record)}
				</Space>
			)
		}
	]

	const renderDeleteApp = (record: TApplicationItem) => {
		return (
			<Popconfirm
				title="Delete this App?"
				onConfirm={async () => {
					const res = await deleteApplication({ id: record?.id })

					return afterModalformFinish(actionRef, res)
				}}
			>
				<Typography.Link type="danger">Delete</Typography.Link>
			</Popconfirm>
		)
	}

	const renderSwitch = (record: TApplicationItem) => {
		return (
			<Switch
				unCheckedChildren="OFF"
				checkedChildren="ON"
				checked={record?.status === GLOBAL_STATUS.ON}
				onChange={async () => {
					const res = await toggleAppStatus({ id: record?.id })

					return afterModalformFinish(actionRef, res)
				}}
			/>
		)
	}

	const renderAddEditApp = (type: 'ADD' | 'EDIT', record?: TApplicationItem) => {
		const isEdit = type === 'EDIT'

		return (
			<ModalForm
				{...MODAL_FORM_PROPS}
				initialValues={isEdit ? record : {}}
				title={isEdit ? 'Edit Application' : 'Add Application'}
				trigger={isEdit ? <Typography.Link>Edit</Typography.Link> : <Button type="primary">ADD</Button>}
				onOpenChange={visible => {
					if (!visible) {
						setImgUrl('')
					} else {
						setImgUrl(record?.image || '')
					}
				}}
				onFinish={async params => {
					let res
					const payload = omit({ ...params, image: imgUrl }, ['logo'])

					if (isEdit) {
						res = await updateApplication({ ...payload, id: record?.id })
					} else {
						res = await addApplication(payload)
					}

					if (!res?.data.success){
						return false
					}

					return afterModalformFinish(actionRef, res)
				}}
			>
				<ProFormText label="Name" name="name" rules={[{ required: true }]} />
				<ProFormUploadButton
					label="Image"
					name="logo"
					fieldProps={{
						accept: 'image/*',
						listType: 'picture-card',
						fileList: imgUrl ? [{ uid: '-1', name: 'image.png', status: 'done', url: imgUrl }] : [],
						beforeUpload: BeforeUpload,
						multiple: false,
						maxCount: 1,
						customRequest: async e => {
							const res = await uploadImage(e?.file)

							setImgUrl(res?.data.data?.url)
						},
						onRemove: () => setImgUrl('')
					}}
				/>
			</ModalForm>
		)
	}

	const fetchData = async () => {
		const res = await getApplication({})

		return {
			data: res?.data.data
		}
	}

	return (
		<div>
			<ProTable
				{...PRO_TABLE_PROPS}
				rowKey="id"
				columns={columns}
				actionRef={actionRef}
				request={fetchData}
				toolBarRender={() => [renderAddEditApp('ADD')]}
				scroll={{ x: 500 }}
			/>
		</div>
	)
}

export default Application
