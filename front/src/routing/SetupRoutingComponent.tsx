import React, {useState} from "react";
import {Router, SetupRouter} from "./router";
import {RouterProvider} from "react-router-dom";
import {GetStatus} from "../api/api";
import {IsSetupRequired} from "../api/setup";

interface IProps {
}

interface IState {
    isLoading?: boolean;
    setupRequired?: boolean;
}

export default class SetupRoutingComponent extends React.Component<IProps, IState> {
    constructor(props: IProps) {
        super(props);

        this.state = {
            isLoading: true,
            setupRequired: true,
        };
    }

    componentDidMount() {
        IsSetupRequired().then((response) => {
            this.setState({
                isLoading: false,
                setupRequired: response,
            })
        });
    }

    componentDidUpdate() {
        console.log(this.state);
    }

    router() {
        if (this.state.isLoading) {
            return <div>Loading...</div>
        }

        if (this.state.setupRequired) {
           return (<RouterProvider router={SetupRouter} />)
        } else {
           return (<RouterProvider router={Router} />)
        }
    }

    render() {
        return (
            <div>
                {this.router()}
            </div>
        )
    }
}