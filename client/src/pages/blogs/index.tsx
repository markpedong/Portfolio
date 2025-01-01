import { TBlogsItem, addBlogs, deleteBlogs, getBlogs, toggleBlogStatus, updateBlogs, uploadImage } from '@/api'
import { GLOBAL_STATUS } from '@/api/constants'
import { INPUT_LINK, MODAL_FORM_PROPS, PRO_TABLE_PROPS } from '@/constants'
import { INPUT_TRIM, dateTimeFormatter } from '@/utils'
import { BeforeUpload, afterModalformFinish } from '@/utils/antd'
import {
	ActionType,
	ModalForm,
	ProColumns,
	ProFormDatePicker,
	ProFormText,
	ProFormTextArea,
	ProFormUploadButton,
	ProTable
} from '@ant-design/pro-components'
import { Button, Image, Popconfirm, Space, Switch, Typography } from 'antd'
import { useRef, useState } from 'react'

const Blogs = () => {
	const [imgUrl, setImgUrl] = useState('')
	const actionRef = useRef<ActionType>()
	const columns: ProColumns<TBlogsItem>[] = [
		{
			title: 'Image',
			align: 'center',
			width: 80,
			search: false,
			render: (_, record) => <Image src={record?.image} alt="image" width={60} height={60} />
		},
		{
			title: 'Title',
			align: 'start',
			dataIndex: 'title'
		},
		{
			title: 'Description',
			align: 'center',
			dataIndex: 'description',
			ellipsis: true,
			width: 200
		},
		{
			title: 'Link',
			align: 'center',
			dataIndex: 'link',
			ellipsis: true,
			width: 200
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
			width: 160,
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
			width: 160,
			render: (_, record) => (
				<Space>
					{renderSwitch(record)}
					{renderAddEditBlogs('EDIT', record)}
					{renderDeleteBlogs(record)}
				</Space>
			)
		}
	]

	const renderSwitch = (record: TBlogsItem) => {
		return (
			<Switch
				unCheckedChildren="OFF"
				checkedChildren="ON"
				checked={record?.status === GLOBAL_STATUS.ON}
				onChange={async () => {
					const res = await toggleBlogStatus({ id: record?.id })

					return afterModalformFinish(actionRef, res)
				}}
			/>
		)
	}

	const renderAddEditBlogs = (type: 'ADD' | 'EDIT', record?: TBlogsItem) => {
		const isEdit = type === 'EDIT'
		return (
			<ModalForm
				{...MODAL_FORM_PROPS}
				initialValues={isEdit ? record : {}}
				title={isEdit ? 'Edit Blogs' : 'Add Blogs'}
				trigger={
					isEdit ? (
						<Typography.Link onClick={() => setImgUrl(record?.image!)}>Edit</Typography.Link>
					) : (
						<Button type="primary">ADD</Button>
					)
				}
				grid
				onOpenChange={visible => !visible && setImgUrl('')}
				layout="inline"
				rowProps={{
					gutter: [0, 25]
				}}
				onFinish={async params => {
					let res

					if (isEdit) {
						res = await updateBlogs({ ...params, image: imgUrl, id: record?.id })
					} else {
						res = await addBlogs({ ...params, image: imgUrl })
					}

					return afterModalformFinish(actionRef, res)
				}}
			>
				<ProFormText label="Title" name="title" rules={[{ required: true }]} />
				<ProFormText {...INPUT_TRIM} label="Link" name="link" rules={[INPUT_LINK, { required: true }]} />
				<ProFormDatePicker rules={[{ required: true }]} label="Date Written" name="date" />
				<ProFormTextArea label="Description" name="description" rules={[{ required: true }]} />
				<ProFormUploadButton
					label="Image"
					name="image"
					rules={[{ required: true }]}
					fieldProps={{
						accept: 'image/*',
						listType: 'picture-card',
						fileList: imgUrl ? [{ uid: '-1', name: 'image.png', status: 'done', url: imgUrl }] : [],
						beforeUpload: BeforeUpload,
						multiple: false,
						maxCount: 1,
						customRequest: async e => {
							const res = await uploadImage(e?.file)

							setImgUrl(res?.data?.data.url)
						},
						onRemove: () => setImgUrl('')
					}}
				/>
			</ModalForm>
		)
	}

	const renderDeleteBlogs = (record: TBlogsItem) => {
		return (
			<Popconfirm
				title="Delete this Blog?"
				onConfirm={async () => {
					const res = await deleteBlogs({ id: record?.id })

					return afterModalformFinish(actionRef, res)
				}}
			>
				<Typography.Link type="danger">Delete</Typography.Link>
			</Popconfirm>
		)
	}

	const fetchData = async () => {
		const res = await getBlogs()

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
				toolBarRender={() => [renderAddEditBlogs('ADD')]}
				scroll={{ x: 1100 }}
			/>
		</div>
	)
}

export default Blogs
