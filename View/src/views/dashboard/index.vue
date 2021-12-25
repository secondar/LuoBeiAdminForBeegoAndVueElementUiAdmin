<template>
  <div class="dashboard-container">
    <el-container>
      <el-header class="header">
        <el-row>
          <el-button
            @click="windowOpwnUrl('https://doc.imoecg.com')"
            type="primary"
            >前端文档</el-button
          >
          <el-button
            @click="windowOpwnUrl('https://doc.imoecg.com')"
            type="success"
            >GO后端文档</el-button
          >
          <el-button
            @click="windowOpwnUrl('https://doc.imoecg.com')"
            type="info"
            >PHP后端文档</el-button
          >
          <el-button
            @click="windowOpwnUrl('https://doc.imoecg.com')"
            type="warning"
            >JAVA后端文档</el-button
          >
          <el-button
            @click="windowOpwnUrl('https://github.com/secondar/LuoBeiAdminForBeegoAndVueElementUiAdmin.git')"
            type="danger"
            >Github</el-button
          >
          <el-button @click="windowOpwnUrl('https://gitee.com/secondar/luo-bei-admin-for-beego-and-vue-element-ui-admin.git')"
            >Gitee</el-button
          >
        </el-row>
      </el-header>
      <el-main>
        <div class="message">
          <div class="title">萝北后台管理系统</div>
          <div class="content">
            一直想做一款后台管理系统，看了很多优秀的开源项目但是发现没有合适自己的。<br />
            因为自己不是很喜欢用已经完完整整的较为臃肿的后台管理系统，仅仅就想要一个做好权限管理就好了的后台管理系统<br />
            于是利用空闲休息时间开始自己写一套后台系统。
            如此有了萝北管理系统。<br />
            它可以用于所有的Web应用程序，如网站管理后台，网站会员中心，CMS，CRM，OA等等<br />
            当然，您也可以对它进行深度定制，以做出更强系统。<br />
            所有前端后台代码封装过后十分精简易上手，出错概率低。<br />
            同时支持移动客户端访问。系统会陆续更新一些实用功能。
          </div>
        </div>
        <div class="info">
          <el-row :type="'flex'" justify="space-between">
            <el-col class="" :span="7">
              <el-card>
                <div slot="header">
                  <span class="title">联系信息</span>
                </div>
                <div class="item">
                  <i class="el-icon-s-promotion"></i>
                  官网：
                  <el-link
                    href="https://www.imoecg.com"
                    type="info"
                    target="_blank"
                    >www.imoecg.com</el-link
                  >
                </div>
                <div class="item">
                  <i class="el-icon-s-promotion"></i>
                  博客：
                  <el-link
                    href="https://www.bugquit.com"
                    type="info"
                    target="_blank"
                    >www.bugquit.com</el-link
                  >
                </div>
                <div class="item">
                  <svg-icon icon-class="example" />
                  邮箱：
                  <el-link type="info">mail@bugquit.com</el-link>
                </div>
              </el-card>
            </el-col>
            <el-col class="" :span="7">
              <el-card>
                <div slot="header">
                  <span class="title">更新日志</span>
                </div>
                <el-collapse>
                  <el-collapse-item title="v1.0.0  - 2021-12-25">
                    <div>1.萝北前后端分离后台管理系统正式发布</div>
                  </el-collapse-item>
                </el-collapse>
              </el-card>
            </el-col>
            <el-col class="" :span="9">
              <el-card>
                <div slot="header">
                  <span class="title">捐赠支持</span>
                </div>
                <div>
                  <div class="qrcode">
                    <img
                      src="https://www.bugquit.com/wp-content/uploads/2019/12/logo-1.png"
                      alt=""
                    />
                    <img
                      src="https://www.bugquit.com/wp-content/uploads/2019/12/logo-1.png"
                      alt=""
                    />
                  </div>
                  <div class="tip">您可以请作者喝杯雪王表示鼓励</div>
                </div>
              </el-card>
            </el-col>
          </el-row>
        </div>
      </el-main>
    </el-container>
  </div>
</template>

<script>
export default {
  name: "Dashboard",
  data() {
    return {};
  },
  computed: {},
  created() {},
  methods: {
    windowOpwnUrl(url) {
      window.open(url);
    },
  },
};
</script>
<style>
.dashboard-container {
  padding-top: 3rem;
  padding-left: 1rem;
}
.header {
  text-align: center;
}
.message {
  color: rgb(125, 125, 126);
}
.message .title {
  font-size: 1.5rem;
  margin-bottom: 1rem;
}
.info {
  margin-top: 5rem;
}
.info .title {
  color: rgb(125, 125, 126);
}
.info .item {
  margin-bottom: 0.5rem;
  font-size: 0.8rem;
  color: rgb(125, 125, 126);
}
.tip {
  color: rgb(125, 125, 126);
  font-size: 0.8rem;
}
.qrcode img {
  width: 50%;
}
</style>