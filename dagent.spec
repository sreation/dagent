#软件包简要介绍
Summary: build dagent
#软件包的名字
Name: dagent
#软件包的主版本号
Version: 1.0
#软件包的次版本号
Release: 0
#源代码包，默认将在上面提到的SOURCES目录中寻找
Source0: %{name}-%{version}.tar.gz
#授权协议
License: GPL
#软件分类
Group: Development/Tools
#软件包的内容介绍
%description
 
#表示预操作字段，后面的命令将在源码代码BUILD前执行
%prep
 
#BUILD字段，将通过直接调用源码目录中自动构建工具完成源码编译操作
%build
 
#file
#安装字段
%install
 
# 二进制执行文件
mkdir -p ${RPM_BUILD_ROOT}/usr/local/bin
cp -f /data/dagent/dagent ${RPM_BUILD_ROOT}/usr/local/bin/dagent
mkdir -p ${RPM_BUILD_ROOT}/etc/dagent
cp -f /data/dagent/config.xml ${RPM_BUILD_ROOT}/etc/dagent/config.xml
mkdir -p ${RPM_BUILD_ROOT}/etc/init.d
cp -f /data/dagent/dagent_inid_script ${RPM_BUILD_ROOT}/etc/init.d/dagent_inid_script
 
%files
 
%defattr(-,root,root)
 
/usr/local/bin/dagent
/etc/dagent/config.xml
/etc/init.d/dagent_inid_script 
%dir
 
#/etc/transcoding
