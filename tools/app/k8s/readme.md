#### 1. Mac install Minikube

1. 安装 Home-brew https://cloud.tencent.com/developer/article/1853162

   ```shell
   > /usr/bin/ruby -e "$(curl -fsSL https://cdn.jsdelivr.net/gh/ineo6/homebrew-install/install)"
   ```

2. 安装 Minukube https://minikube.sigs.k8s.io/docs/start/

   ```shell
   > brew install minikube
   # 遇到问题: No such file or directory @ rb_sysopen - ...
   # 解决方式：https://zhuanlan.zhihu.com/p/491515480
   ```

3. ```shell
   > kubectl version -o json # 显示版本信息
   ```

#### 2. Centos7.9 install kubeadmin

1. 查看版本号
   ```shell
   >  cat /etc/redhat-release
   ```
2. 添加 IP
   ```shell
   > ip add
   ```
3. 修改主机名称
   ```shell
   > hostnamectl set-hostname k8s-master && bash
   ```
4. 添加 hosts,这里的 IP 是自己服务器 ip
   ```shell
   > ifconfig #查看主机IP地址
     cat >> /etc/hosts << EOF
     x.x.x.x  k8s-master
     EOF
   ```
5. 关闭防火墙,关闭 selinux
   ```shell
   > systemctl stop firewalld
   > systemctl disable firewalld
   > sed -i 's/enforcing/disabled/' /etc/selinux/config # 永久
   > setenforce 0 # 临时
   ```
6. 关闭 swap
   ```shell
   > swapoff -a # 临时
   > sed -i 's/.*swap.*/#&/' /etc/fstab # 永久
   ```
7. 将桥接的 IPv4 流量传递到 iptables 的链

   ```shell
   > cat > /etc/sysctl.d/k8s.conf << EOF
     net.bridge.bridge-nf-call-ip6tables = 1
     net.bridge.bridge-nf-call-iptables = 1
     EOF

   > sysctl --system # 生效
   ```

8. 时间同步
   ```shell
   > yum install ntpdate -y
   > ntpdate time.windows.com
   ```
9. 安装 Docker

   ```shell
   > wget https://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo -O /etc/yum.repos.d/docker-ce.repo
   > yum -y install docker-ce-20.10.12-3.el7
   > systemctl enable docker && systemctl start docker && systemctl status docker
   > docker --version

   ```

10. 给 docker 添加加速器

    ```shell
    > cat > /etc/docker/daemon.json << EOF
    {
      "registry-mirrors": ["https://qj799ren.mirror.aliyuncs.com"],
      "exec-opts": ["native.cgroupdriver=systemd"],
      "log-driver": "json-file",
      "log-opts": {
      "max-size": "100m"
    },
      "storage-driver": "overlay2"
    }
    EOF

    > systemctl restart docker

    ```

11. 添加 kubernetes 的 yum 源

    ```shell
    > cat > /etc/yum.repos.d/kubernetes.repo << EOF
      [kubernetes]
      name=Kubernetes
      baseurl=https://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-x86_64
      enabled=1
      gpgcheck=0
      repo_gpgcheck=0
      gpgkey=https://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg
      https://mirrors.aliyun.com/kubernetes/yum/doc/rpm-package-key.gpg
      EOF

    ```

12. 安装 kubeadm，kubelet 和 kubectl
    ```shell
    > yum -y install kubelet-1.21.5-0 kubeadm-1.21.5-0 kubectl-1.21.5-0  #当前时间最新版是v1.21.5固定版本，下面有用
    > systemctl enable kubelet
    ```
13. 部署 Kubernetes Master

    ```shell
    > kubeadm init --apiserver-advertise-address=10.0.4.11  --image-repository registry.aliyuncs.com/google_containers  --kubernetes-version v1.21.5  --service-cidr=10.96.0.0/12  --pod-network-cidr=10.244.0.0/16 --ignore-preflight-errors=all

    ```

    - 参数说明：
      - –apiserver-advertise-address=10.0.4.11 这个参数就是 master 主机的 IP 地址，例如我的 Master 主机的 IP 是：10.0.4.11
      - –image-repository registry.aliyuncs.com/google_containers 这个是镜像地址，由于国外地址无法访问，故使用的阿里云仓库地址：repository registry.aliyuncs.com/google_containers
      - –kubernetes-version=v1.21.5 这个参数是下载的 k8s 软件版本号
      - –service-cidr=10.96.0.0/12 这个参数后的 IP 地址直接就套用 10.96.0.0/12 ,以后安装时也套用即可，不要更改
      - –pod-network-cidr=10.244.0.0/16 k8s 内部的 pod 节点之间网络可以使用的 IP 段，不能和 service-cidr 写一样，如果不知道怎么配，就先用这个 10.244.0.0/16
      - –ignore-preflight-errors=all 添加这个会忽略错误
    - 执行语句后，看到如下的信息说明就安装成功了。
      ![img.png](img.png)  


14. 执行如下语句
    ```shell
    > mkdir -p $HOME/.kube
    > sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
    > sudo chown $(id -u):$(id -g) $HOME/.kube/config
    > kubectl get nodes    #节点状态为NotReady
    ```
15. 安装 Pod 网络插件（CNI）
    ```shell
    > wget https://docs.projectcalico.org/archive/v3.20/manifests/calico.yaml
    ```

```
---
# Source: calico/templates/calico-config.yaml
# This ConfigMap is used to configure a self-hosted Calico installation.
kind: ConfigMap
apiVersion: v1
metadata:
  name: calico-config
  namespace: kube-system
data:
  # Typha is disabled.
  typha_service_name: &#34;none&#34;
  # Configure the backend to use.
  calico_backend: &#34;bird&#34;

  # Configure the MTU to use
  veth_mtu: &#34;1440&#34;

  # The CNI network configuration to install on each node.  The special
  # values in this config will be automatically populated.
  cni_network_config: |-
    {
      &#34;name&#34;: &#34;k8s-pod-network&#34;,
      &#34;cniVersion&#34;: &#34;0.3.1&#34;,
      &#34;plugins&#34;: [
        {
          &#34;type&#34;: &#34;calico&#34;,
          &#34;log_level&#34;: &#34;info&#34;,
          &#34;datastore_type&#34;: &#34;kubernetes&#34;,
          &#34;nodename&#34;: &#34;__KUBERNETES_NODE_NAME__&#34;,
          &#34;mtu&#34;: __CNI_MTU__,
          &#34;ipam&#34;: {
              &#34;type&#34;: &#34;calico-ipam&#34;
          },
          &#34;policy&#34;: {
              &#34;type&#34;: &#34;k8s&#34;
          },
          &#34;kubernetes&#34;: {
              &#34;kubeconfig&#34;: &#34;__KUBECONFIG_FILEPATH__&#34;
          }
        },
        {
          &#34;type&#34;: &#34;portmap&#34;,
          &#34;snat&#34;: true,
          &#34;capabilities&#34;: {&#34;portMappings&#34;: true}
        }
      ]
    }

---
# Source: calico/templates/kdd-crds.yaml
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: felixconfigurations.crd.projectcalico.org
spec:
  scope: Cluster
  group: crd.projectcalico.org
  version: v1
  names:
    kind: FelixConfiguration
    plural: felixconfigurations
    singular: felixconfiguration
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: ipamblocks.crd.projectcalico.org
spec:
  scope: Cluster
  group: crd.projectcalico.org
  version: v1
  names:
    kind: IPAMBlock
    plural: ipamblocks
    singular: ipamblock

---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: blockaffinities.crd.projectcalico.org
spec:
  scope: Cluster
  group: crd.projectcalico.org
  version: v1
  names:
    kind: BlockAffinity
    plural: blockaffinities
    singular: blockaffinity

---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: ipamhandles.crd.projectcalico.org
spec:
  scope: Cluster
  group: crd.projectcalico.org
  version: v1
  names:
    kind: IPAMHandle
    plural: ipamhandles
    singular: ipamhandle

---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: ipamconfigs.crd.projectcalico.org
spec:
  scope: Cluster
  group: crd.projectcalico.org
  version: v1
  names:
    kind: IPAMConfig
    plural: ipamconfigs
    singular: ipamconfig

---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: bgppeers.crd.projectcalico.org
spec:
  scope: Cluster
  group: crd.projectcalico.org
  version: v1
  names:
    kind: BGPPeer
    plural: bgppeers
    singular: bgppeer

---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: bgpconfigurations.crd.projectcalico.org
spec:
  scope: Cluster
  group: crd.projectcalico.org
  version: v1
  names:
    kind: BGPConfiguration
    plural: bgpconfigurations
    singular: bgpconfiguration

---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: ippools.crd.projectcalico.org
spec:
  scope: Cluster
  group: crd.projectcalico.org
  version: v1
  names:
    kind: IPPool
    plural: ippools
    singular: ippool

---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: hostendpoints.crd.projectcalico.org
spec:
  scope: Cluster
  group: crd.projectcalico.org
  version: v1
  names:
    kind: HostEndpoint
    plural: hostendpoints
    singular: hostendpoint

---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: clusterinformations.crd.projectcalico.org
spec:
  scope: Cluster
  group: crd.projectcalico.org
  version: v1
  names:
    kind: ClusterInformation
    plural: clusterinformations
    singular: clusterinformation

---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: globalnetworkpolicies.crd.projectcalico.org
spec:
  scope: Cluster
  group: crd.projectcalico.org
  version: v1
  names:
    kind: GlobalNetworkPolicy
    plural: globalnetworkpolicies
    singular: globalnetworkpolicy

---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: globalnetworksets.crd.projectcalico.org
spec:
  scope: Cluster
  group: crd.projectcalico.org
  version: v1
  names:
    kind: GlobalNetworkSet
    plural: globalnetworksets
    singular: globalnetworkset

---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: networkpolicies.crd.projectcalico.org
spec:
  scope: Namespaced
  group: crd.projectcalico.org
  version: v1
  names:
    kind: NetworkPolicy
    plural: networkpolicies
    singular: networkpolicy

---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: networksets.crd.projectcalico.org
spec:
  scope: Namespaced
  group: crd.projectcalico.org
  version: v1
  names:
    kind: NetworkSet
    plural: networksets
    singular: networkset
---
# Source: calico/templates/rbac.yaml

# Include a clusterrole for the kube-controllers component,
# and bind it to the calico-kube-controllers serviceaccount.
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: calico-kube-controllers
rules:
  # Nodes are watched to monitor for deletions.
  - apiGroups: [&#34;&#34;]
    resources:
      - nodes
      verbs:
        - watch
        - list
        - get
      # Pods are queried to check for existence.
      - apiGroups: [&#34;&#34;]
        resources:
          - pods
          verbs:
            - get
          # IPAM resources are manipulated when nodes are deleted.
          - apiGroups: [&#34;crd.projectcalico.org&#34;]
            resources:
              - ippools
              verbs:
                - list
              - apiGroups: [&#34;crd.projectcalico.org&#34;]
                resources:
                  - blockaffinities
                  - ipamblocks
                  - ipamhandles
                  verbs:
                    - get
                    - list
                    - create
                    - update
                    - delete
                  # Needs access to update clusterinformations.
                  - apiGroups: [&#34;crd.projectcalico.org&#34;]
                    resources:
                      - clusterinformations
                      verbs:
                        - get
                        - create
                        - update
---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: calico-kube-controllers
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: calico-kube-controllers
subjects:
  - kind: ServiceAccount
    name: calico-kube-controllers
    namespace: kube-system
---
# Include a clusterrole for the calico-node DaemonSet,
# and bind it to the calico-node serviceaccount.
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: calico-node
rules:
  # The CNI plugin needs to get pods, nodes, and namespaces.
  - apiGroups: [&#34;&#34;]
    resources:
      - pods
      - nodes
      - namespaces
      verbs:
        - get
      - apiGroups: [&#34;&#34;]
        resources:
          - endpoints
          - services
          verbs:
            # Used to discover service IPs for advertisement.
            - watch
            - list
            # Used to discover Typhas.
            - get
          - apiGroups: [&#34;&#34;]
            resources:
              - nodes/status
              verbs:
                # Needed for clearing NodeNetworkUnavailable flag.
                - patch
                # Calico stores some configuration information in node annotations.
                - update
              # Watch for changes to Kubernetes NetworkPolicies.
              - apiGroups: [&#34;networking.k8s.io&#34;]
                resources:
                  - networkpolicies
                  verbs:
                    - watch
                    - list
                  # Used by Calico for policy information.
                  - apiGroups: [&#34;&#34;]
                    resources:
                      - pods
                      - namespaces
                      - serviceaccounts
                      verbs:
                        - list
                        - watch
                      # The CNI plugin patches pods/status.
                      - apiGroups: [&#34;&#34;]
                        resources:
                          - pods/status
                          verbs:
                            - patch
                          # Calico monitors various CRDs for config.
                          - apiGroups: [&#34;crd.projectcalico.org&#34;]
                            resources:
                              - globalfelixconfigs
                              - felixconfigurations
                              - bgppeers
                              - globalbgpconfigs
                              - bgpconfigurations
                              - ippools
                              - ipamblocks
                              - globalnetworkpolicies
                              - globalnetworksets
                              - networkpolicies
                              - networksets
                              - clusterinformations
                              - hostendpoints
                              - blockaffinities
                              verbs:
                                - get
                                - list
                                - watch
                              # Calico must create and update some CRDs on startup.
                              - apiGroups: [&#34;crd.projectcalico.org&#34;]
                                resources:
                                  - ippools
                                  - felixconfigurations
                                  - clusterinformations
                                  verbs:
                                    - create
                                    - update
                                  # Calico stores some configuration information on the node.
                                  - apiGroups: [&#34;&#34;]
                                    resources:
                                      - nodes
                                      verbs:
                                        - get
                                        - list
                                        - watch
                                      # These permissions are only requried for upgrade from v2.6, and can
                                      # be removed after upgrade or on fresh installations.
                                      - apiGroups: [&#34;crd.projectcalico.org&#34;]
                                        resources:
                                          - bgpconfigurations
                                          - bgppeers
                                          verbs:
                                            - create
                                            - update
                                          # These permissions are required for Calico CNI to perform IPAM allocations.
                                          - apiGroups: [&#34;crd.projectcalico.org&#34;]
                                            resources:
                                              - blockaffinities
                                              - ipamblocks
                                              - ipamhandles
                                              verbs:
                                                - get
                                                - list
                                                - create
                                                - update
                                                - delete
                                              - apiGroups: [&#34;crd.projectcalico.org&#34;]
                                                resources:
                                                  - ipamconfigs
                                                  verbs:
                                                    - get
                                                  # Block affinities must also be watchable by confd for route aggregation.
                                                  - apiGroups: [&#34;crd.projectcalico.org&#34;]
                                                    resources:
                                                      - blockaffinities
                                                      verbs:
                                                        - watch
                                                      # The Calico IPAM migration needs to get daemonsets. These permissions can be
                                                      # removed if not upgrading from an installation using host-local IPAM.
                                                      - apiGroups: [&#34;apps&#34;]
                                                        resources:
                                                          - daemonsets
                                                          verbs:
                                                            - get
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: calico-node
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: calico-node
subjects:
  - kind: ServiceAccount
    name: calico-node
    namespace: kube-system

---
# Source: calico/templates/calico-node.yaml
# This manifest installs the calico-node container, as well
# as the CNI plugins and network config on
# each master and worker node in a Kubernetes cluster.
kind: DaemonSet
apiVersion: apps/v1
metadata:
  name: calico-node
  namespace: kube-system
  labels:
    k8s-app: calico-node
spec:
  selector:
    matchLabels:
      k8s-app: calico-node
  updateStrategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 1
  template:
    metadata:
      labels:
        k8s-app: calico-node
      annotations:
        # This, along with the CriticalAddonsOnly toleration below,
        # marks the pod as a critical add-on, ensuring it gets
        # priority scheduling and that its resources are reserved
        # if it ever gets evicted.
        scheduler.alpha.kubernetes.io/critical-pod: &#39;&#39;
    spec:
      nodeSelector:
        beta.kubernetes.io/os: linux
      hostNetwork: true
      tolerations:
        # Make sure calico-node gets scheduled on all nodes.
        - effect: NoSchedule
          operator: Exists
        # Mark the pod as a critical add-on for rescheduling.
        - key: CriticalAddonsOnly
          operator: Exists
        - effect: NoExecute
          operator: Exists
      serviceAccountName: calico-node
      # Minimize downtime during a rolling upgrade or deletion; tell Kubernetes to do a &#34;force
      # deletion&#34;: https://kubernetes.io/docs/concepts/workloads/pods/pod/#termination-of-pods.
      terminationGracePeriodSeconds: 0
      priorityClassName: system-node-critical
      initContainers:
        # This container performs upgrade from host-local IPAM to calico-ipam.
        # It can be deleted if this is a fresh installation, or if you have already
        # upgraded to use calico-ipam.
        - name: upgrade-ipam
          image: calico/cni:v3.11.3
          command: [&#34;/opt/cni/bin/calico-ipam&#34;, &#34;-upgrade&#34;]
            env:
              - name: KUBERNETES_NODE_NAME
                valueFrom:
                  fieldRef:
                    fieldPath: spec.nodeName
              - name: CALICO_NETWORKING_BACKEND
                valueFrom:
                  configMapKeyRef:
                    name: calico-config
                    key: calico_backend
              volumeMounts:
                - mountPath: /var/lib/cni/networks
                  name: host-local-net-dir
                - mountPath: /host/opt/cni/bin
                  name: cni-bin-dir
              securityContext:
                privileged: true
              # This container installs the CNI binaries
              # and CNI network config file on each node.
              - name: install-cni
                image: calico/cni:v3.11.3
                command: [&#34;/install-cni.sh&#34;]
                  env:
                    # Name of the CNI config file to create.
                    - name: CNI_CONF_NAME
                      value: &#34;10-calico.conflist&#34;
                    # The CNI network config to install on each node.
                    - name: CNI_NETWORK_CONFIG
                      valueFrom:
                        configMapKeyRef:
                          name: calico-config
                          key: cni_network_config
                    # Set the hostname based on the k8s node name.
                    - name: KUBERNETES_NODE_NAME
                      valueFrom:
                        fieldRef:
                          fieldPath: spec.nodeName
                    # CNI MTU Config variable
                    - name: CNI_MTU
                      valueFrom:
                        configMapKeyRef:
                          name: calico-config
                          key: veth_mtu
                    # Prevents the container from sleeping forever.
                    - name: SLEEP
                      value: &#34;false&#34;
                    volumeMounts:
                      - mountPath: /host/opt/cni/bin
                        name: cni-bin-dir
                      - mountPath: /host/etc/cni/net.d
                        name: cni-net-dir
                    securityContext:
                      privileged: true
                    # Adds a Flex Volume Driver that creates a per-pod Unix Domain Socket to allow Dikastes
                    # to communicate with Felix over the Policy Sync API.
                    - name: flexvol-driver
                      image: calico/pod2daemon-flexvol:v3.11.3
                      volumeMounts:
                        - name: flexvol-driver-host
                          mountPath: /host/driver
                      securityContext:
                        privileged: true
                    containers:
                      # Runs calico-node container on each Kubernetes node.  This
                      # container programs network policy and routes on each
                      # host.
                      - name: calico-node
                        image: calico/node:v3.11.3
                        env:
                          # Use Kubernetes API as the backing datastore.
                          - name: DATASTORE_TYPE
                            value: &#34;kubernetes&#34;
                          # Wait for the datastore.
                          - name: WAIT_FOR_DATASTORE
                            value: &#34;true&#34;
                          # Set based on the k8s node name.
                          - name: NODENAME
                            valueFrom:
                              fieldRef:
                                fieldPath: spec.nodeName
                          # Choose the backend to use.
                          - name: CALICO_NETWORKING_BACKEND
                            valueFrom:
                              configMapKeyRef:
                                name: calico-config
                                key: calico_backend
                          # Cluster type to identify the deployment type
                          - name: CLUSTER_TYPE
                            value: &#34;k8s,bgp&#34;
                          # Auto-detect the BGP IP address.
                          - name: IP
                            value: &#34;autodetect&#34;
                          # Enable IPIP
                          - name: CALICO_IPV4POOL_IPIP
                            value: &#34;Always&#34;
                          # Set MTU for tunnel device used if ipip is enabled
                          - name: FELIX_IPINIPMTU
                            valueFrom:
                              configMapKeyRef:
                                name: calico-config
                                key: veth_mtu
                          # The default IPv4 pool to create on startup if none exists. Pod IPs will be
                          # chosen from this range. Changing this value after installation will have
                          # no effect. This should fall within &#96;--cluster-cidr&#96;.
                          - name: CALICO_IPV4POOL_CIDR
                            value: &#34;10.244.0.0/16&#34;
                          # Disable file logging so &#96;kubectl logs&#96; works.
                          - name: CALICO_DISABLE_FILE_LOGGING
                            value: &#34;true&#34;
                          # Set Felix endpoint to host default action to ACCEPT.
                          - name: FELIX_DEFAULTENDPOINTTOHOSTACTION
                            value: &#34;ACCEPT&#34;
                          # Disable IPv6 on Kubernetes.
                          - name: FELIX_IPV6SUPPORT
                            value: &#34;false&#34;
                          # Set Felix logging to &#34;info&#34;
                          - name: FELIX_LOGSEVERITYSCREEN
                            value: &#34;info&#34;
                          - name: FELIX_HEALTHENABLED
                            value: &#34;true&#34;
                        securityContext:
                          privileged: true
                        resources:
                          requests:
                            cpu: 250m
                        livenessProbe:
                          exec:
                            command:
                              - /bin/calico-node
                              - -felix-live
                              - -bird-live
                          periodSeconds: 10
                          initialDelaySeconds: 10
                          failureThreshold: 6
                        readinessProbe:
                          exec:
                            command:
                              - /bin/calico-node
                              - -felix-ready
                              - -bird-ready
                          periodSeconds: 10
                        volumeMounts:
                          - mountPath: /lib/modules
                            name: lib-modules
                            readOnly: true
                          - mountPath: /run/xtables.lock
                            name: xtables-lock
                            readOnly: false
                          - mountPath: /var/run/calico
                            name: var-run-calico
                            readOnly: false
                          - mountPath: /var/lib/calico
                            name: var-lib-calico
                            readOnly: false
                          - name: policysync
                            mountPath: /var/run/nodeagent
                    volumes:
                      # Used by calico-node.
                      - name: lib-modules
                        hostPath:
                          path: /lib/modules
                      - name: var-run-calico
                        hostPath:
                          path: /var/run/calico
                      - name: var-lib-calico
                        hostPath:
                          path: /var/lib/calico
                      - name: xtables-lock
                        hostPath:
                          path: /run/xtables.lock
                          type: FileOrCreate
                      # Used to install CNI.
                      - name: cni-bin-dir
                        hostPath:
                          path: /opt/cni/bin
                      - name: cni-net-dir
                        hostPath:
                          path: /etc/cni/net.d
                      # Mount in the directory for host-local IPAM allocations. This is
                      # used when upgrading from host-local to calico-ipam, and can be removed
                      # if not using the upgrade-ipam init container.
                      - name: host-local-net-dir
                        hostPath:
                          path: /var/lib/cni/networks
                      # Used to create per-pod Unix Domain Sockets
                      - name: policysync
                        hostPath:
                          type: DirectoryOrCreate
                          path: /var/run/nodeagent
                      # Used to install Flex Volume Driver
                      - name: flexvol-driver-host
                        hostPath:
                          type: DirectoryOrCreate
                          path: /usr/libexec/kubernetes/kubelet-plugins/volume/exec/nodeagent~uds
---

apiVersion: v1
kind: ServiceAccount
metadata:
  name: calico-node
  namespace: kube-system

---
# Source: calico/templates/calico-kube-controllers.yaml

# See https://github.com/projectcalico/kube-controllers
apiVersion: apps/v1
kind: Deployment
metadata:
  name: calico-kube-controllers
  namespace: kube-system
  labels:
    k8s-app: calico-kube-controllers
spec:
  # The controllers can only have a single active instance.
  replicas: 1
  selector:
    matchLabels:
      k8s-app: calico-kube-controllers
  strategy:
    type: Recreate
  template:
    metadata:
      name: calico-kube-controllers
      namespace: kube-system
      labels:
        k8s-app: calico-kube-controllers
      annotations:
        scheduler.alpha.kubernetes.io/critical-pod: &#39;&#39;
    spec:
      nodeSelector:
        beta.kubernetes.io/os: linux
      tolerations:
        # Mark the pod as a critical add-on for rescheduling.
        - key: CriticalAddonsOnly
          operator: Exists
        - key: node-role.kubernetes.io/master
          effect: NoSchedule
      serviceAccountName: calico-kube-controllers
      priorityClassName: system-cluster-critical
      containers:
        - name: calico-kube-controllers
          image: calico/kube-controllers:v3.11.3
          env:
            # Choose which controllers to run.
            - name: ENABLED_CONTROLLERS
              value: node
            - name: DATASTORE_TYPE
              value: kubernetes
          readinessProbe:
            exec:
              command:
                - /usr/bin/check-status
                - -r
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: calico-kube-controllers
  namespace: kube-system
---
# Source: calico/templates/calico-etcd-secrets.yaml
---
# Source: calico/templates/calico-typha.yaml
---
# Source: calico/templates/configure-canal.yaml
EOF
```

注意唯一需要修改的是 CALICO_IPV4POOL_CIDR 对应的 IP，需要与前面 kubeadm init 的 --pod-network-cidr 指定的一样。–pod-network-cidr=10.244.0.0/16

```shell
> kubectl apply -f calico.yaml
```

16. 验证网络
    ```shell
    > kubectl get nodes  #看到 Ready说明网络就正常了
    > kubectl get pods -n kube-system  #全部显示Running就是对的
    ```
17. 单集版的 k8s 安装后, 无法部署服务
    ```shell
    # 因为默认master不能部署pod,有污点, 需要去掉污点或者新增一个node，我这里是去除污点。
    > kubectl get node -o yaml | grep taint -A 5 #执行后看到有输出说明有污点
    > kubectl taint nodes --all node-role.kubernetes.io/master-   #执行这句就行，就是取消污点
    ```
18. 安装补全命令的包
    ```shell
    > yum -y install bash-completion  #安装补全命令的包
    > kubectl completion bash
    > source /usr/share/bash-completion/bash_completion
    > kubectl completion bash >/etc/profile.d/kubectl.sh
    > source /etc/profile.d/kubectl.sh
    > cat  >>  /root/.bashrc <<EOF
    source /etc/profile.d/kubectl.sh
    EOF
    ```
19. 测试 kubernetes 集群
    ```shell
    # 在Kubernetes集群中部署一个Nginx：
    > kubectl create deployment nginx --image=nginx
    > kubectl expose deployment nginx --port=80 --type=NodePort
    > kubectl get pods,svc
    # 注意：看到对外暴露的是3xxxx端口,开放安全组端口范围
    ```

## ps:

centos 安装 k8s 基本参考 【kubernetes 最新版安装单机版 v1.21.5】，之所以 cpoy 一遍文章内容，主要担心原文丢失！！

参考链接：

- [kubernetes 最新版安装单机版 v1.21.5](https://blog.csdn.net/qq_14910065/article/details/122180162)
- [安装 Pod 网络插件（CNI）](https://blog.csdn.net/moxiaotang/article/details/124790965)

[应用集合](../readme.md)
