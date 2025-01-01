import {
	TPortfolioItem,
	addPortfolios,
	deletePortfolios,
	getPortfolios,
	togglePortfolioStatus,
	updatePortfolios,
	uploadImage
} from '@/api'
import { GLOBAL_STATUS } from '@/api/constants'
import { INPUT_LINK, MODAL_FORM_PROPS, PRO_TABLE_PROPS } from '@/constants'
import { capFrstLtr, randomColorGenerator } from '@/constants/helper'
import { INPUT_TRIM, dateTimeFormatter } from '@/utils'
import { BeforeUpload, afterModalformFinish } from '@/utils/antd'
import {
	ActionType,
	ModalForm,
	ProColumns,
	ProFormSelect,
	ProFormText,
	ProFormUploadButton,
	ProTable
} from '@ant-design/pro-components'
import { Button, Image, Popconfirm, Space, Switch, Tag, Typography } from 'antd'
import { useRef, useState } from 'react'

const Portfolio = () => {
	const [imgUrl, setImgUrl] = useState('')
	const actionRef = useRef<ActionType>()
	const columns: ProColumns<TPortfolioItem>[] = [
		{
			title: 'Image',
			align: 'center',
			search: false,
			width: 100,
			render: (_, record) => <Image src={record?.image} alt="logo" width={60} height={60} />
		},
		{
			title: 'Title',
			align: 'start',
			dataIndex: 'title'
		},
		{
			title: 'Tech Stack',
			align: 'left',
			search: false,
			render: (_, record) =>
				record?.tech?.map(q => (
					<Tag color={`#${randomColorGenerator()}`} key={q}>
						{capFrstLtr(q)}
					</Tag>
				))
		},
		{
			title: 'Link',
			align: 'center',
			dataIndex: 'link',
			ellipsis: true,
			width: 225
		},
		{
			title: (
				<div className="flex flex-col gap-0">
					<div>Created</div>
					<div>Updated</div>
				</div>
			),
			align: 'center',
			width: 160,
			search: false,
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
			width: 160,
			search: false,
			render: (_, record) => {
				return (
					<Space>
						{renderSwitch(record)}
						{renderAddEditPortfolio('EDIT', record)}
						{renderDeletePortfolio(record)}
					</Space>
				)
			}
		}
	]

	const renderSwitch = (record: TPortfolioItem) => {
		return (
			<Switch
				unCheckedChildren="OFF"
				checkedChildren="ON"
				checked={record?.status === GLOBAL_STATUS.ON}
				onChange={async () => {
					const res = await togglePortfolioStatus({ id: record?.id })

					return afterModalformFinish(actionRef, res)
				}}
			/>
		)
	}

	const renderAddEditPortfolio = (type: 'ADD' | 'EDIT', record?: TPortfolioItem) => {
		const isEdit = type === 'EDIT'
		return (
			<ModalForm
				{...MODAL_FORM_PROPS}
				title={isEdit ? 'Edit Portfolio' : 'Add Portfolio'}
				trigger={
					isEdit ? (
						<Typography.Link onClick={() => setImgUrl(record?.image!)}>Edit</Typography.Link>
					) : (
						<Button type="primary">Add</Button>
					)
				}
				width={1000}
				initialValues={isEdit ? record : {}}
				onOpenChange={visible => !visible && setImgUrl('')}
				labelCol={{ flex: '110px' }}
				onFinish={async params => {
					let res

					if (isEdit) {
						res = await updatePortfolios({ ...params, image: imgUrl, id: record?.id })
					} else {
						res = await addPortfolios({ ...params, image: imgUrl })
					}

					return afterModalformFinish(actionRef, res)
				}}
			>
				<ProFormText label="Title" name="title" rules={[{ required: true }]} />
				<ProFormText {...INPUT_TRIM} label="Link" name="link" rules={[INPUT_LINK, { required: true }]} />
				<ProFormSelect
					label="Tech Stack"
					name="tech"
					fieldProps={{ mode: 'tags', maxCount: 7 }}
					rules={[{ required: true }]}
				/>
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

	const renderDeletePortfolio = (record: TPortfolioItem) => {
		return (
			<Popconfirm
				title="Delete this Portfolio?"
				onConfirm={async () => {
					const res = await deletePortfolios({ id: record?.id })

					return afterModalformFinish(actionRef, res)
				}}
			>
				<Typography.Link type="danger">Delete</Typography.Link>
			</Popconfirm>
		)
	}

	const fetchData = async () => {
		const res = await getPortfolios()

		return {
			data: res?.data.data
		}
	}

	return (
		<div>
			<ProTable
				{...PRO_TABLE_PROPS}
				rowKey="id"
				actionRef={actionRef}
				columns={columns}
				request={fetchData}
				toolBarRender={() => [renderAddEditPortfolio('ADD')]}
				scroll={{ x: 1100 }}
			/>
		</div>
	)
}

export default Portfolio
