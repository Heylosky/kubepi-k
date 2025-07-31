import Layout from "@/business/app-layout/horizontal-layout"

const Tektons = {
    path: "/tektons",
    sort: 5,
    component: Layout,
    requirePermission: {
        resource: "clusters",
        verb: "list"
    },
    meta: {
        title: "business.tekton.list",
        icon: "iconfont iconkubernets",
    },
    children: [
        {
            path: "",
            component: () => import("@/business/tekton-management"),
            name: "Tektons",
            meta: {
                title: "business.tekton.list",
                activeMenu: "/tektons",
            }
        }
    ]
}

export default Tektons