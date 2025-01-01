import {
	TTestimonialsItem,
	addTestimonials,
	deleteTestimonials,
	getTestimonials,
	toggleTestimonialStatus,
	updateTestimonials,
	uploadImage
} from '@/api'
import { GLOBAL_STATUS } from '@/api/constants'
import { MODAL_FORM_PROPS, PRO_TABLE_PROPS } from '@/constants'
import { dateTimeFormatter } from '@/utils'
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

const Testimonials = () => {
	const [imgUrl, setImgUrl] = useState('')
	const actionRef = useRef<ActionType>()
	const columns: ProColumns<TTestimonialsItem>[] = [
		{
			title: 'Image',
			align: 'center',
			search: false,
			render: (_, record) => <Image src={record?.image} alt="image" width={60} height={60} />
		},
		{
			title: 'Author',
			align: 'center',
			dataIndex: 'author'
		},
		{
			title: 'Description',
			align: 'center',
			dataIndex: 'description',
			width: 200,
			ellipsis: true
		},
		{
			title: 'Job',
			align: 'center',
			dataIndex: 'job'
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
			width: 180,
			search: false,
			render: (_, record) => (
				<Space>
					{renderSwitch(record)}
					{renderAddEditTestimonials('EDIT', record)}
					{renderDeleteTestimonials(record)}
				</Space>
			)
		}
	]

	const renderSwitch = (record: TTestimonialsItem) => {
		return (
			<Switch
				unCheckedChildren="OFF"
				checkedChildren="ON"
				checked={record?.status === GLOBAL_STATUS.ON}
				onChange={async () => {
					const res = await toggleTestimonialStatus({ id: record?.id })

					return afterModalformFinish(actionRef, res)
				}}
			/>
		)
	}

	const renderAddEditTestimonials = (type: 'ADD' | 'EDIT', record?: TTestimonialsItem) => {
		const isEdit = type === 'EDIT'
		return (
			<ModalForm
				{...MODAL_FORM_PROPS}
				initialValues={isEdit ? record : {}}
				title={isEdit ? 'Edit Testimonials' : 'Add Testimonials'}
				trigger={
					isEdit ? (
						<Typography.Link onClick={() => setImgUrl(record?.image!)}>Edit</Typography.Link>
					) : (
						<Button type="primary">ADD</Button>
					)
				}
				grid
				layout="inline"
				rowProps={{
					gutter: [0, 25]
				}}
				onOpenChange={(visible) => !visible && setImgUrl("")}
				onFinish={async params => {
					let res

					if (isEdit) {
						res = await updateTestimonials({ ...params, image: imgUrl, id: record?.id })
					} else {
						res = await addTestimonials({ ...params, image: imgUrl })
					}

					return afterModalformFinish(actionRef, res)
				}}
			>
				<ProFormText label="Author" name="author" rules={[{ required: true }]} />
				<ProFormText label="Job" name="job" rules={[{ required: true }]} />
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

							setImgUrl(res?.data.data?.url)
						},
						onRemove: () => setImgUrl('')
					}}
				/>
			</ModalForm>
		)
	}
	const renderDeleteTestimonials = (record: TTestimonialsItem) => {
		return (
			<Popconfirm
				title="Delete this Testimonial?"
				onConfirm={async () => {
					const res = await deleteTestimonials({ id: record?.id })

					return afterModalformFinish(actionRef, res)
				}}
			>
				<Typography.Link type="danger">Delete</Typography.Link>
			</Popconfirm>
		)
	}

	const fetchData = async () => {
		const res = await getTestimonials()

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
				toolBarRender={() => [renderAddEditTestimonials('ADD')]}
				scroll={{ x: 950 }}
			/>
		</div>
	)
}

export default Testimonials
