import Layout from "@/business/app-layout/horizontal-layout"

const Dmps = {
    path: "/dmps",
    sort: 5,
    component: Layout,
    requirePermission: {
        resource: "clusters",
        verb: "list"
    },
    meta: {
        title: "business.dmp.list",
        icon: "iconfont iconkubernets",
    },
    children: [
        {
            path: "",
            component: () => import("@/business/dmp-management"),
            name: "Dmps",
            meta: {
                title: "business.dmp.list",
                activeMenu: "/dmps",
            }
        }
    ]
}

export default Dmps