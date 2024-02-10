import * as React from "react";
import * as ReactDOM from "react-dom/client";
import {
    createBrowserRouter,
    RouterProvider,
} from "react-router-dom";
import "../index.css";
import Error404 from "../pages/error/404";
import CurrentStep from "../pages/setup/CurrentStep";
import CreateAdminUserPage from "../pages/setup/step/CreateAdminUserPage";
import ImportKubernetesClusterPage from "../pages/setup/step/ImportKubernetesClusterPage";
import Portal from "../pages/portal/Portal";

export const SetupRouter  = createBrowserRouter([
    {
        path: "/",
        children: [
            {
                path: "setup",
                children: [
                    {
                        path: "CreateAdminUser",
                        element: <CreateAdminUserPage/>,
                    },
                    {
                        path: "ImportKubernetesCluster",
                        element: <ImportKubernetesClusterPage/>,
                    },
                    {
                        path: "",
                        element: <CurrentStep/>,
                    }
                    ]
            },
            {
                path: "",
                element: <CurrentStep/>,
            }
        ]
    },
    {
        path: "*",
        element: <Error404/>,
    }
])

export const Router = createBrowserRouter([
    {
        path: "/",
        element: <Portal/>,
    },
]);
