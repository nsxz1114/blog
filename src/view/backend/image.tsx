﻿import { Button, Image, Modal, Space, Table, Upload, message } from "antd";
import { useEffect, useState } from "react";
import { imageDelete, imageList, imageType, imageUpload } from "../../api/image";
import type { ColumnsType } from "antd/es/table";
import { UploadOutlined } from "@ant-design/icons";
import type { RcFile } from "antd/es/upload/interface";
import { paramsType } from "@/api";

interface PaginationState extends paramsType {
    total: number;
}

export const AdminImage = () => {
    const [loading, setLoading] = useState(false);
    const [data, setData] = useState<imageType[]>([]);
    const [pagination, setPagination] = useState<PaginationState>({
        page: 1,
        page_size: 10,
        total: 0,
    });

    // 获取图片列表
    const fetchData = async (page = 1) => {
        try {
            setLoading(true);

            const res = await imageList({
                page,
                page_size: pagination.page_size,
            });
            if (res.code === 2000) {
                setData(res.data.list);
                setPagination({
                    ...pagination,
                    page,
                    total: res.data.total,
                });
            } else {
                message.error(res.msg);
            }
        } catch (error) {
            console.error("获取图片列表失败:", error);
            message.error("获取图片列表失败");
        } finally {
            setLoading(false);
        }
    };

    // 删除图片
    const handleDelete = async (id: number) => {
        Modal.confirm({
            title: "确认删除",
            content: "确定要删除这张图片吗？删除后不可恢复。",
            onOk: async () => {
                try {
                    const res = await imageDelete(id);
                    if (res.code === 2000) {
                        message.success("删除成功");
                        fetchData(pagination.page);
                    } else {
                        message.error(res.msg);
                    }
                } catch (error) {
                    message.error("删除失败");
                    console.error("删除失败:", error);
                }
            },
        });
    };

    // 上传图片
    const handleUpload = async (file: RcFile) => {
        try {
            const files = file instanceof FileList ? Array.from(file) : [file];
            const res = await imageUpload(files);
            if (res.code === 2000) {
                message.success("上传成功");
                fetchData(pagination.page);
            } else {
                message.error(res.msg);
            }
        } catch (error) {
            console.error("上传失败:", error);
            message.error("上传失败");
        }
        return false;
    };

    // 添加文件类型和大小限制
    const uploadProps = {
        beforeUpload: handleUpload,
        showUploadList: false,
        accept: 'image/*',
        maxSize: 20 * 1024 * 1024, // 20MB
        multiple: true,
    };

    const columns: ColumnsType<imageType> = [
        {
            title: "ID",
            dataIndex: "id",
            key: "id",
        },
        {
            title: "预览",
            key: "preview",
            render: (_, record) => (
                <Image
                    src={record.path}
                    alt={record.name}
                    width={100}
                    height={100}
                    style={{ objectFit: "cover" }}
                />
            ),
        },
        {
            title: "文件名",
            dataIndex: "name",
            key: "name",
        },
        {
            title: "类型",
            dataIndex: "type",
            key: "type",
        },
        {
            title: "大小",
            dataIndex: "size",
            key: "size",
            render: (size) => `${(size / 1024).toFixed(2)} KB`,
        },
        {
            title: "上传时间",
            dataIndex: "created_at",
            key: "created_at",
        },
        {
            title: "操作",
            key: "action",
            render: (_, record) => (
                <Space size="middle">
                    <Button
                        type="link"
                        danger
                        onClick={() => handleDelete(record.id)}
                    >
                        删除
                    </Button>
                </Space>
            ),
        },
    ];

    useEffect(() => {
        fetchData();
    }, []);

    return (
        <div className="admin_image">
            <div>
                <Upload.Dragger {...uploadProps}>
                    <p className="ant-upload-drag-icon">
                        <UploadOutlined style={{ fontSize: 48, color: '#40a9ff' }} />
                    </p>
                    <p className="ant-upload-text">点击或拖拽图片到此区域上传</p>
                    <p className="ant-upload-hint">
                        支持同时上传多张图片，单个文件大小不超过20MB
                    </p>
                </Upload.Dragger>
            </div>
            <Table
                columns={columns}
                dataSource={data}
                rowKey="id"
                pagination={{
                    ...pagination,
                    position: ['bottomCenter'],
                    style: { marginTop: '16px' }
                }}
                loading={loading}
                onChange={(pagination) => fetchData(pagination.current)}
            />
        </div>
    );
};