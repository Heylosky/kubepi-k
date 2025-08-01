import el from "element-ui/lib/locale/lang/en"
import fu from "fit2cloud-ui/src/locale/lang/en_US" // 加载fit2cloud的内容


const message = {
    commons: {
        message_box: {
            alert: "Alter",
            confirm: "Confirm",
            prompt: "Prompt",
        },
        personal: {
            profile: "Profile",
            exit: "Logout",
            project_url: "Project Address",
            issue: "Feedback",
            talk: "Participate in the discussion",
            star: "Star",
            version: "Version",
            copy_right: "Copyright © 2021-2024 FIT2CLOUD",
            introduction: "KubePi is a modern K8s panel.",
            introduction2: "KubePi allows administrators to import multiple Kubernetes clusters and assign permissions of different clusters and namespaces to specific users through permission control. It allows developers to manage and troubleshoot applications running in the Kubernetes cluster to better deal with the complexity in the Kubernetes cluster."
        },
        help: {
            help: "Help",
            about: "About",
            ko_docs: "Official Document",
            dev_doc: "Development Document",
        },
        button: {
            detail: "Detail",
            delete: "Delete",
            skip: "Skip",
            import: "Import",
            create: "Create",
            cancel: "Cancel",
            login: "Login",
            confirm: "Confirm",
            add: "Add",
            edit: "Edit",
            all_select: "All",
            upload: "Upload",
            search: "Search",
            rbac_manage: "RBAC Manage",
            sync: "Sync",
            bind: "Bind",
            change_password: 'Change Password',
            export: 'Export all helm release',
        },
        table: {
            name: "Name",
            kind: "Kind",
            created_time: "Created Time",
            age: "Age",
            status: "Status",
            action: "Action",
            creat_by: "created by",
            built_in: "Built in",
            description: "Description",
            empty_text: "There are no rows to show.",
            mfa_enable: "MFA Auth"
        },
        header: {
            help_doc: "document",
            support: "support",
            guide: "guide",
            guide_text: "Thank you for choosing this product. Would you like to register your first cluster now?"
        },
        bool: {
            true: "true",
            false: "false"
        },
        enable: {
          true: "Enable",
          false: "Disable"
        },
        search: {
            quickSearch: "Search"
        },
        form: {
            select_placeholder: "please select"
        },
        validate: {
            limit: "长度在 {0} 到 {1} 个字符",
            input: "please input {0}",
            select: "please select {0}",
            required: "required",
            email: "please input a valid email",
            number_limit: "Please enter the correct number",
            password_help: "Valid password: 8-30 digits, English letters + numbers + special characters (~!@_#$%^&*)",
            name_not_compliant: "The name does not conform to the naming convention!",
            name_rules: "Only lowercase English, numbers and-are supported",
        },
        msg: {
            create_success: "create success",
            delete_success: "delete success",
            update_success: "update success",
            duplicate_failed: "The operation failed because the data is duplicate！",
            no_data: "no data",
        },
        confirm_message: {
            delete: "This operation cannot be undone. Do you want to continue?",
            create_success: "create success",
            save_success: "save success",
        },
        login: {
            username_or_email: "username or email",
            password: "password",
            title: "login KubePi",
            welcome: "Welcome back, please enter your user name and password to log in",
            expires: "The authentication information has expired. Please log in again.",
            mfa_helper: "Scan the QR code below with the MFA Authenticator app to get a 6-digit verification code",
            mfa_login_helper:"Please enter 6-digit verification for MFA Authenticator",
        },
        sso: {
            title: "SSO Redirecting",
        },
    },
    business: {
        cluster: {
            cluster: "Cluster",
            namespace: "Namespace",
            scope: "Scope",
            version: "Version",
            list: "Clusters",
            import: "import cluster",
            edit_cluster: "edit cluster",
            edit: "edit",
            nodes: "Nodes",
            label: "Label",
            description: "Description",
            cluster_detail: "Detail",
            connect_setting: "Connect setting",
            connect_direction: "Connect direction",
            connect_forward: "Forward",
            connect_backward: "Backward",
            authenticate_setting: "Authenticate setting",
            certificate: "Certificate",
            authenticate_mode: "Authenticate mode",
            expect: "expect",
            management: "management",
            open_dashboard: "Console",
            cluster_version: "version",
            member: "Member",
            role: "Role",
            administrator: "administrator",
            viewer: "viewer",
            custom: "custom",
            contains_namespace: "(contains namespace roles)",
            rule: "rule",
            config_file: "kubeconfig file",
            config_content: "Config content",
            hidden_cluster: "hide inaccessible clusters",
            namespace_role_form_check_msg: "namespace or role list cannot be empty",
            api_group: "API Groups",
            resource: "resources",
            verb: "verbs",
            cluster_role_form_check_msg: "api groups,resources or verbs  list cannot be empty",
            user_not_in_cluster: "You are not a member of this cluster and cannot access the console of this cluster. Please contact the administrator to add you as a cluster member.",
            ready: "Ready",
            not_ready: "NotReady",
            repo: "Repo",
            repo_auth: "Repo Authorization"
        },
        cluster_role: {
            none: "None",
            cluster_administrator: "Administrator",
            cluster_viewer: "Viewer",
        },
        user: {
            user_management: "Users",
            username: "username",
            nickname: "Nickname",
            email: "Email",
            user_list: "User list",
            role_list: "Role list",
            user: "User",
            role: "Role",
            template: "template",
            base_on_exists_role: "base on exists role",
            permission: "permission",
            permission_setting: "Permission setting",
            password: "password",
            confirm_password: "confirm password",
            old_password: "old password",
            new_password: "new password",
            change_password: "change password",
            resource_name: "resource name",
            please_input_password: "please input password",
            please_input_password_agin: "please input password again",
            password_not_equal: "two passwords are inconsistent",
            ldap: "LDAP",
            ldap_address: "Address",
            ldap_tls: "tls",
            ldap_port: "Port",
            ldap_username: "Username",
            ldap_password: "Password",
            ldap_filter_dn: "User Filtering DN",
            ldap_filter_rule: "User Filtering Rules",
            ldap_helper: "Note: Users who cannot get the Name mapping attribute will not be matched",
            ldap_sync: "Start syncing, please check the user list later",
            ldap_sync_error: "Please save first",
            type: "Type",
            ldap_mapping: "User Attribute Mapping",
            ldap_mapping_helper: "User attribute mapping represents how to map user attributes in LDAP to kubepi users, name, nickName, email are the attributes required by kubepi users",
            ldap_test: "Test connection",
            sso: "SSO",
            sso_protocol: "Protocol",
            sso_interface_address: "Interface Address",
            sso_client_id: "Client ID",
            sso_client_secret: "Client Secret",
            sso_x509_cert: "Certificate",
            sso_x509_key: "Certificate Key",
            sso_idp_metadata_url: "IDP Metadata URL",
            sso_test_result: "Test SSO connection successful",
            sso_helper: "Note: If SSO is enabled, if you want to log in with a local account, please directly access \"http(s)://host/kubepi/login\"",
            sso_remake: "Reset",
            sso_test: "Test connection",
            test_result: "The connection is successful, matching {count} users",
            test_login: "Test Login",
            login_success: "Test login success",
            login_failed: "Login failed",
            import_user: "Import User",
            import_enable: "Whether it can be imported",
            import_user_success: "Import successful",
            import_user_failed: "Import failed user {user}",
            ldap_remake: "Reset",
            time_limit: "Connection timeout",
            size_limit: "Number of search pages",
        },
        system: {
            system_log: "Auditing Log",
            operation_log: "Operation Logs",
            operator: "Operator",
            operation: "Operation",
            operation_domain: "Resource",
            specific_information: "Informations",
            login_log: "Login Logs",
            username: "Username",
            ip: "Login ip",
            city: "Login city",
        },
        image_repos: {
          list: "Image Registries",
          name: "Name",
          endpoint: "Address",
          downloadUrl: "Download URL",
          username: "Username",
          password: "Password",
          type: "Type",
          repo: "Mirror Library",
          load_repo: "Load image library",
          auth: "Auth",
          allow_anonymous: "Allow anonymous docker pull ",
          repo_null: "Repo is null",
          images: "Image List",
          push_image: "Push the mirror to the current project",
          version: "Version",
        },
        dmp: {
            list: "DMP List",
            name: "Name",
            label: "Label",
            version: "Version",
            open_dmp: "Console",
            user_not_in_cluster: "You are not a member of this cluster and cannot access the DMP console of this cluster. Please contact the administrator to add you as a cluster member.",
            ready: "Ready",
            not_ready: "NotReady",
        },
        tekton: {
            list: "Tekton List",
            name: "Name",
            label: "Label",
            version: "Version",
            open_tekton: "Console",
            user_not_in_cluster: "You are not a member of this cluster and cannot access the Tekton console of this cluster. Please contact the administrator to add you as a cluster member.",
            ready: "Ready",
            not_ready: "NotReady",
        },
    },
}

const description = {
    i18n_user_administrator: "Super administrator, with permissions for all objects.",
    i18n_user_manage_cluster: "The Cluster Administrator has all the permissions of the cluster object.",
    i18n_user_manage_rbac: "Role and user administrators have all the permissions of the user objects.",
    i18n_user_manage_repo: "Image repostries administrators have all the permissions of the image repostries objects.",
    i18n_user_manage_readonly: "Read only user with access to all objects only,",
    i18n_user_common_user: "Ordinary users only have access to cluster objects",
    i18n_user_manage_chart: "Chart warehouse administrator, has all rights to the Chart warehouse",

    i18n_cluster_owner: "The cluster owner has permissions on all objects",
    i18n_cluster_viewer: "A cluster read-only user who has the read-only permission on all objects",
    i18n_manage_cluster_rbac: "Cluster Access control Administrator who has all permissions on ClusterRole and ClusterRoleBinding objects",
    i18n_view_cluster_rbac: "Cluster access control read-only user who has the ClusterRole and ClusterRoleBinding read-only permission",
    i18n_manage_cluster_storage: "Cluster storage administrator who has all permissions on the StorageClass and PersistentVolume objects",
    i18n_view_cluster_storage: "The cluster stores read-only users and has the read-only permission on the StorageClass and PersistentVolume objects",
    i18n_manage_namespaces: "Namespace administrator who has all permissions on Namespace objects",
    i18n_view_namespaces: "A read-only Namespace user who has all permissions on Namespace objects",
    i18n_view_events: "A cluster event read-only user who has read-only permission on Events objects",
    i18n_view_nodes: "A read-only user who has the read-only permission on Node objects",

    i18n_manage_nodes: "The node administrator has read-only permissions on node objects",
    i18n_manage_crd: "The custom resource administrator has all the permissions of the CustomResourceDefinition objec",
    i18n_view_crd: "User defined resource administrator with read-only permission on CustomResourceDefinition object",

    i18n_manage_config: "The configuration administrator has all permissions on the current namespace configmap, secret, resourcequotes, limitranges, horizontalpodautoscalers and poddisruptionbudget objects",
    i18n_view_config: "Configure a read-only user with read-only permissions on the current namespace configmap, secret, resourcequotes, limitranges, horizontalpodautoscalers and poddisruptionbudget objects",
    i18n_namespace_owner: "The namespace owner has all permissions on all objects in the current namespace",
    i18n_namespace_viewer: "Namespace read-only user with read-only permission for all objects in the current namespace",
    i18n_view_workload: "Workload read-only user with read-only permissions for daemonset, statefulset, deployment, job, cronjob and pod in the current namespace",
    i18n_manage_workload: "Workload administrator, read-only user of workload, with all permissions of daemonset, statefulset, deployment, job, cronjob and pod in the current namespace",
    i18n_manage_storage: "Storage administrator who has all permissions on the persistentvolumeclaim object in the current namespace",
    i18n_view_storage: "Stores a read-only user with read-only permissions on the persistentvolumeclaim object in the current namespace",
    i18n_view_service_discovery: "The service found a read-only user with read-only permissions on service, endpoint, progress and networkpolicy objects in the current namespace",
    i18n_manage_service_discovery: "The service discovery administrator has all permissions on service, endpoint, ingress and networkpolicy objects in the current namespace",
    i18n_manage_rbac: "The service discovery administrator has all permissions on service, endpoint, ingress and networkpolicy objects in the current namespace",
    i18n_view_rbac: "Namespace access control read-only user with read-only permissions for role, rolebinding and serviceaccount objects in the current namespace",
    i18n_manage_appmarket: "Application market administrator, who has all rights to the application market"
}

const apiObjects = {
    users: "users",
    roles: "roles",
    clusters: "clusters",
    systems: "systems",
}

const apiVerbs = {
    "update": "update",
    "delete": "delete",
    "get": "get",
    "list": "list",
    "create": "create",
    "authorization": "authorization"
}

const system_logs = {
    post: "create",
    put: "update",
    delete: "delete",
    clusters: "Clusters",
    users: "User",
    roles: "Role",
    systems: "Auditing Log",
    clusters_members: "Cluster Member",
    clusters_clusterroles: "Cluster Role",
    clusters_repos: "Cluster Repos",
    imagerepos: "Image Registries",
    ldap: "LDAP",
}


export default {
    ...el,
    ...fu,
    ...message,
    ...apiObjects,
    ...apiVerbs,
    ...description,
    ...system_logs
}
