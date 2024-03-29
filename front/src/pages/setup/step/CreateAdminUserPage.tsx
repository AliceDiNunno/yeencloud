import React from "react";
import MailInput from "../../../components/inputs/MailInput";
import PasswordInput from "../../../components/inputs/PasswordInput";
import FormPageComponent from "../../../components/FormPage";

export default class CreateAdminUserPage extends React.Component {
    render() {
        return (
            <FormPageComponent>
                <h1>Create Admin User</h1>

                <div>
                    <MailInput value={""} onChange={(value) => console.log(value)}/>
                    <PasswordInput value={""} onChange={(value) => console.log(value)}/>
                </div>
            </FormPageComponent>
        )
    }
}