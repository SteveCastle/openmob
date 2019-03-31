import React from 'react'
import PropTypes from 'prop-types'
import Content from '@openmob/bluebird/src/components/layout/Content'
import Card from '@openmob/bluebird/src/components/cards/Card'
import Form from '@openmob/bluebird/src/components/forms/Form'
import Widget from '@openmob/bluebird/src/components/forms/Widget'
import Label from '@openmob/bluebird/src/components/forms/Label'
import Input from '@openmob/bluebird/src/components/forms/Input'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const CreateContact = ({ id }) => (
  <Content>
    <Card>
      <Form>
        <h1>Create Contact</h1>
        <Widget>
          <Label>FirstName</Label>
          <Input placeholder="String" />
        </Widget>
        <Widget>
          <Label>MiddleName</Label>
          <Input placeholder="String" />
        </Widget>
        <Widget>
          <Label>LastName</Label>
          <Input placeholder="String" />
        </Widget>
        <Widget>
          <Label>Email</Label>
          <Input placeholder="String" />
        </Widget>
        <Widget>
          <Label>PhoneNumber</Label>
          <Input placeholder="String" />
        </Widget>

        <Button label="Create" block variant="primary" />
      </Form>
    </Card>
  </Content>
)

CreateContact.propTypes = {
  id: PropTypes.string,
}

export default CreateContact
