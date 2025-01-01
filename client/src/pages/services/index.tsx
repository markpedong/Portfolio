import {
	TServiceItem,
	addServices,
	deleteServices,
	getServices,
	updateServices,
	toggleServiceStatus,
	uploadImage
} from '@/api'
import { GLOBAL_STATUS } from '@/api/constants'
import { MODAL_FORM_PROPS, PRO_TABLE_PROPS } from '@/constants'
import { INPUT_TRIM, dateTimeFormatter } from '@/utils'
import { BeforeUpload, afterModalformFinish } from '@/utils/antd'
import {
	ActionType,
	ModalForm,
	ProColumns,
	ProFormText,
	ProFormTextArea,
	ProFormUploadButton,
	ProTable
} from '@ant-design/pro-components'
import { Button, Image, Popconfirm, Space, Switch, Typography } from 'antd'
import { useRef, useState } from 'react'

const Services = () => {
	const [imgUrl, setImgUrl] = useState('')
	const actionRef = useRef<ActionType>()
	const columns: ProColumns<TServiceItem>[] = [
		{
			title: 'Logo',
			align: 'center',
			search: false,
			render: (_, record) => <Image src={record?.logo} alt="logo" width={30} height={30} />
		},
		{
			title: 'Title',
			dataIndex: 'title',
			align: 'left',
			width: 200
		},
		{
			title: 'Description',
			dataIndex: 'description',
			align: 'justify',
		},
		{
			title: (
				<div className="flex flex-col gap-0">
					<div>Created</div>
					<div>Updated</div>
				</div>
			),
			width: 180,
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
			width: 190,
			render: (_, record) => {
				return (
					<Space>
						{renderSwitch(record)}
						{renderAddEditService('EDIT', record)}
						{renderDeleteLink(record)}
					</Space>
				)
			}
		}
	]

	const renderSwitch = (record: TServiceItem) => {
		return (
			<Switch
				unCheckedChildren="OFF"
				checkedChildren="ON"
				checked={record?.status === GLOBAL_STATUS.ON}
				onChange={async () => {
					const res = await toggleServiceStatus({ id: record?.id })

					return afterModalformFinish(actionRef, res)
				}}
			/>
		)
	}

	const renderDeleteLink = (record: TServiceItem) => {
		return (
			<Popconfirm
				title="Delete this Service?"
				onConfirm={async () => {
					const res = await deleteServices({ id: record?.id })

					return afterModalformFinish(actionRef, res)
				}}
			>
				<Typography.Link type="danger">Delete</Typography.Link>
			</Popconfirm>
		)
	}

	const renderAddEditService = (type: 'EDIT' | 'ADD', record?: TServiceItem) => {
		const isEdit = type === 'EDIT'

		return (
			<ModalForm
				{...MODAL_FORM_PROPS}
				initialValues={isEdit ? record : {}}
				title={isEdit ? 'Edit Service' : 'Add Service'}
				labelCol={{ flex: '100px' }}
				trigger={
					isEdit ? (
						<Typography.Link onClick={() => setImgUrl(record?.logo!)}>Edit</Typography.Link>
					) : (
						<Button type="primary">Add</Button>
					)
				}
				onOpenChange={(visible) => !visible && setImgUrl("")}
				onFinish={async params => {
					let res

					if (isEdit) {
						res = await updateServices({ ...params, logo: imgUrl, id: record?.id })
					} else {
						res = await addServices({ ...params, logo: imgUrl })
					}

					return afterModalformFinish(actionRef, res)
				}}
			>
				<ProFormText {...INPUT_TRIM} label="Title" name="title" rules={[{ required: true }]} />
				<ProFormTextArea label="Description" name="description" rules={[{ required: true }]} />
				<ProFormUploadButton
					label="Logo"
					name="logo"
					rules={[{ required: true }]}
					fieldProps={{
						accept: 'image/*',
						listType: 'picture-card',
						fileList: imgUrl ? [{ uid: '-1', name: 'image', status: 'done', url: imgUrl }] : [],
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
		const res = await getServices()

		return {
			data: res?.data.data
		}
	}

	return (
		<div>
			<ProTable
				{...PRO_TABLE_PROPS}
				actionRef={actionRef}
				toolBarRender={() => [renderAddEditService('ADD')]}
				rowKey="id"
				columns={columns}
				request={fetchData}
				scroll={{ x: 1000 }}
			/>
		</div>
	)
}

export default Services
