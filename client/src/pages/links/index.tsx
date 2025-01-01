import { TLinksItem, addLinks, deleteLinks, getLinks, toggleLinkStatus, updateLinks } from '@/api'
import { GLOBAL_STATUS } from '@/api/constants'
import { INPUT_LINK, MODAL_FORM_PROPS, PRO_TABLE_PROPS } from '@/constants'
import { INPUT_TRIM, dateTimeFormatter } from '@/utils'
import { afterModalformFinish } from '@/utils/antd'
import { ActionType, ModalForm, ProColumns, ProFormSelect, ProFormText, ProTable } from '@ant-design/pro-components'
import { Button, Popconfirm, Space, Switch, Typography } from 'antd'
import React, { useRef, useState } from 'react'
import { FaGithubAlt, FaInstagram, FaLinkedin, FaTwitter } from 'react-icons/fa6'

const Links = () => {
	const actionRef = useRef<ActionType>()
	const [existingTypes, setExistingTypes] = useState<string[]>([])
	const columns: ProColumns<TLinksItem>[] = [
		{
			title: 'Link',
			dataIndex: 'link',
			align: 'left'
		},
		{
			title: 'Type',
			dataIndex: 'type',
			align: 'left'
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
			render: (_, record) => {
				return (
					<Space align="center">
						{renderSwitch(record)}
						{renderAddEditLinks('EDIT', record)}
						{renderDeleteLink(record)}
					</Space>
				)
			}
		}
	]

	const renderSwitch = (record: TLinksItem) => {
		return (
			<Switch
				unCheckedChildren="OFF"
				checkedChildren="ON"
				checked={record?.status === GLOBAL_STATUS.ON}
				onChange={async () => {
					const res = await toggleLinkStatus({ id: record?.id })

					return afterModalformFinish(actionRef, res)
				}}
			/>
		)
	}

	const renderTitle = (icon: React.ReactNode, title: string) => (
		<Space align="center">
			{icon}
			<span>{title}</span>
		</Space>
	)

	const options = [
		{
			label: renderTitle(<FaGithubAlt />, 'Github'),
			value: 'github'
		},
		{
			label: renderTitle(<FaLinkedin />, 'Linkedin'),
			value: 'linkedin'
		},
		{
			label: renderTitle(<FaInstagram />, 'Instagram'),
			value: 'instagram'
		},
		{
			label: renderTitle(<FaTwitter />, 'Twitter'),
			value: 'twitter'
		}
	].map(item => ({
		...item,
		disabled: existingTypes.includes(item.value)
	}))

	const renderDeleteLink = (q: TLinksItem) => {
		return (
			<Popconfirm
				title="Delete this Link?"
				onConfirm={async () => {
					const res = await deleteLinks({ id: q?.id })

					return afterModalformFinish(actionRef, res)
				}}
			>
				<Typography.Link type="danger">Delete</Typography.Link>
			</Popconfirm>
		)
	}

	const renderAddEditLinks = (type: 'EDIT' | 'ADD', record?: TLinksItem) => {
		const isEdit = type === 'EDIT'
		return (
			<ModalForm
				{...MODAL_FORM_PROPS}
				labelCol={{ flex: '80px' }}
				title={isEdit ? 'Edit Link' : 'Add Link'}
				initialValues={isEdit ? record : {}}
				trigger={
					isEdit ? (
						<Typography.Link>Edit</Typography.Link>
					) : (
						<Button type="primary" disabled={existingTypes?.length === 4}>
							Add Link
						</Button>
					)
				}
				onFinish={async params => {
					let res

					if (isEdit) {
						res = await updateLinks({
							link: params.link,
							type: params.type,
							id: record?.id
						})
					} else {
						res = await addLinks(params)
					}

					return afterModalformFinish(actionRef, res)
				}}
			>
				<ProFormText {...INPUT_TRIM} label="Link" name="link" rules={[INPUT_LINK, { required: true }]} />
				<ProFormSelect
					label="Type"
					name="type"
					rules={[{ required: true }]}
					disabled={existingTypes?.includes(record?.type!)}
					fieldProps={{ options }}
				/>
			</ModalForm>
		)
	}

	const fetchData = async () => {
		const res = await getLinks()

		setExistingTypes(res?.data.data?.map(q => q?.type))
		return {
			data: res?.data?.data ?? []
		}
	}

	return (
		<div>
			<Typography.Title level={5}>Update the links that we are providing on the website</Typography.Title>
			<ProTable
				{...PRO_TABLE_PROPS}
				columns={columns}
				actionRef={actionRef}
				search={false}
				rowKey="id"
				request={fetchData}
				toolBarRender={() => [renderAddEditLinks('ADD')]}
				scroll={{
					x: 800
				}}
			/>
		</div>
	)
}

export default Links
