import React from "react";
import { Routes, Route, Link } from "react-router-dom";
import "../../App.css";
import { HomeOutlined, UsergroupAddOutlined, PercentageOutlined } from "@ant-design/icons";
import { Breadcrumb, Layout, Menu, theme, Button, message } from "antd";
import logo from "../../assets/logo.png";

import HomeCodeandPromotion from "../../pages/admin/palm";
import Codes from "../../pages/admin/palm/code";
import CodeCreate from "../../pages/admin/palm/code/create";
import CodeEdit from "../../pages/admin/palm/code/edit";

import AddAdmin from "../../pages/admin/palm/add";
import AdminCreate from "../../pages/admin/palm/add/create";
import AdminEdit from "../../pages/admin/palm/add/edit";

import Stock2 from "../../pages/stock/Stock2";
import Stock3 from "../../pages/stock/Stock3";
import Stock4 from "../../pages/stock/Stock4";

import AddProduct from "../../pages/product/Create/AddProduct";
import EditProduct from "../../pages/product/Edit/EditProduct";

const { Header, Content, Footer } = Layout;

const AdminLayout: React.FC = () => {

  const page = localStorage.getItem("page");
  const [messageApi, contextHolder] = message.useMessage();
//   const [collapsed, setCollapsed] = useState(false);

  const {
    token: { colorBgContainer },
  } = theme.useToken();

  const setCurrentPage = (val: string) => {
    localStorage.setItem("page", val);
  };

  const Logout = () => {
    localStorage.clear();
    messageApi.success("Logout successful");
    setTimeout(() => {
      location.href = "/";
    }, 2000);
  };

  return (
    <Layout style={{ minHeight: "100vh" }}>
      {contextHolder}

      <Layout>
        {/* Header with horizontal menu */}
        <Header style={{ background: colorBgContainer, display: "flex", justifyContent: "space-between", alignItems: "center" }}>
          <div style={{ display: "flex", alignItems: "center" }}>
            <img src={logo} alt="Logo" style={{ height: 50, marginRight: 16 }} />
            <Menu theme="light" mode="horizontal" defaultSelectedKeys={[page ? page : "dashboard"]}>
              <Menu.Item key="dashboard" onClick={() => setCurrentPage("dashboard")}>
                <Link to="/home">
                  <HomeOutlined />
                  <span>Home</span>
                </Link>
              </Menu.Item>
              <Menu.Item key="code" onClick={() => setCurrentPage("code")}>
                <Link to="/code">
                  <PercentageOutlined />
                  <span>Code</span>
                </Link>
              </Menu.Item>
              <Menu.Item key="add" onClick={() => setCurrentPage("add")}>
                <Link to="/add">
                <UsergroupAddOutlined />
                  <span>AddAdmin</span>
                </Link>
              </Menu.Item>
              <Menu.Item key="Stock2" onClick={() => setCurrentPage("Stock2")}>
                <Link to="/Stock2">
                <UsergroupAddOutlined />
                  <span>Stock2</span>
                </Link>
              </Menu.Item>
              <Menu.Item key="product" onClick={() => setCurrentPage("product")}>
                <Link to="/product">
                <UsergroupAddOutlined />
                  <span>product</span>
                </Link>
              </Menu.Item>
              <Menu.Item key="proedit" onClick={() => setCurrentPage("proedit")}>
                <Link to="/proedit">
                <UsergroupAddOutlined />
                  <span>proedit</span>
                </Link>
              </Menu.Item>
              
            </Menu>
          </div>
          <Button onClick={Logout}>ออกจากระบบ</Button>
        </Header>

        <Content style={{ margin: "0 16px" }}>
          <Breadcrumb style={{ margin: "16px 0" }} />
          <div
            style={{
              padding: 24,
              minHeight: "100%",
              background: colorBgContainer,
            }}
          >
            <Routes>
              <Route path="/home" element={<HomeCodeandPromotion />} />

              <Route path="/code" element={<Codes />} />
              <Route path="/code/create" element={<CodeCreate />} />
              <Route path="/code/edit/:id" element={<CodeEdit />} />

              <Route path="/add" element={<AddAdmin />} />
              <Route path="/add/create" element={<AdminCreate />} />
              <Route path="/add/edit/:id" element={<AdminEdit />} />

              <Route path="/Stock2" element={<Stock2 />} />
              <Route path="/Stock3" element={<Stock3 />} />
              <Route path="/Stock4" element={<Stock4 />} />

              <Route path="/product" element={<AddProduct />} />
              <Route path="/proedit" element={<EditProduct />} />

            </Routes>
          </div>
        </Content>

        <Footer style={{ textAlign: "center" }}>IGOTSOFAR</Footer>
      </Layout>
    </Layout>
  );
};

export default AdminLayout;
