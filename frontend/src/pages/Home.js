import React, { useState } from 'react';
import { Layout, Menu, Dropdown, Button } from 'antd';
import {
  CalendarOutlined,
  TeamOutlined,
  MessageOutlined,
  DownOutlined
} from '@ant-design/icons';

const { Header, Content, Sider } = Layout;

const Home = () => {
  const [collapsed, setCollapsed] = useState(false);

  const onCollapse = () => {
    setCollapsed(!collapsed);
  };

  const menu = (
    <Menu>
      <Menu.Item key="1" >
        Group 1
      </Menu.Item>
      <Menu.Item key="2">
        Group 2
      </Menu.Item>
      <Menu.Item key="3">
        Group 3
      </Menu.Item>
    </Menu>
  );

  return (
    <Layout>
      <Header className="site-layout-background" style={{ padding: 0 }} >
        <Dropdown overlay={menu}>
          <Button style={{
            width: '180px',
            margin: '10px',
            textAlign: 'left',
            background: 'transparent',
            color: '#ffffff',
            border: 'none',
            fontSize: '18px',
            fontWeight: 800,
          }}>
            Group 1 <DownOutlined />
          </Button>
        </Dropdown>
      </Header>
      <Layout style={{ minHeight: '100vh' }}>
        <Sider collapsible style={{ backgroundColor: '#ffffff' }} collapsed={collapsed} onCollapse={onCollapse} >
          <Menu theme="light" defaultSelectedKeys={['1']} mode="inline">
            <Menu.Item key="1" icon={<CalendarOutlined />}>
              Calendar
            </Menu.Item>
            <Menu.Item key="2" icon={<TeamOutlined />}>
              Members
            </Menu.Item>
            <Menu.Item key="3" icon={<MessageOutlined />}>
              ChatRoom
            </Menu.Item>
          </Menu>
        </Sider>
        <Layout className="site-layout">
          <Content style={{ margin: '0 16px' }}>
            <div className="site-layout-background" style={{ padding: 24, minHeight: 360 }}>
              Bill is a cat.
            </div>
          </Content>
        </Layout>
      </Layout>
    </Layout>
  )
}

export default Home;