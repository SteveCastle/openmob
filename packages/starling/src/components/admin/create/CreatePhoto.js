import React from 'react'
import PropTypes from 'prop-types'
import Content from '@openmob/bluebird/src/components/layout/Content'
import Card from '@openmob/bluebird/src/components/cards/Card'
import Form from '@openmob/bluebird/src/components/forms/Form'
import Widget from '@openmob/bluebird/src/components/forms/Widget'
import Label from '@openmob/bluebird/src/components/forms/Label'
import Input from '@openmob/bluebird/src/components/forms/Input'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const CreatePhoto = ({ id }) => (
  <Content>
    <Card>
      <Form>
        <h1>Create Photo</h1>
        <Widget>
          <Label>ID</Label>
          <Input placeholder="ID!" />
        </Widget>
        <Widget>
          <Label>CreatedAt</Label>
          <Input placeholder="Time!" />
        </Widget>
        <Widget>
          <Label>UpdatedAt</Label>
          <Input placeholder="Time!" />
        </Widget>
        <Widget>
          <Label>URI</Label>
          <Input placeholder="String!" />
        </Widget>
        <Widget>
          <Label>Width</Label>
          <Input placeholder="Int!" />
        </Widget>
        <Widget>
          <Label>Height</Label>
          <Input placeholder="Int!" />
        </Widget>

        <Button label="Create" block variant="primary" />
      </Form>
    </Card>
  </Content>
)

CreatePhoto.propTypes = {
  id: PropTypes.string,
}

export default CreatePhoto
