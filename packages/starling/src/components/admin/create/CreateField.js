import React from 'react'
import PropTypes from 'prop-types'
import Input from '@openmob/bluebird/src/components/forms/Input'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const CreateField = ({ id }) => (
  <div>
    <h1>Create Field</h1>
    <div>
      <label>ID</label>
      <Input placeholder="ID!" />
    </div>
    <div>
      <label>CreatedAt</label>
      <Input placeholder="Time!" />
    </div>
    <div>
      <label>UpdatedAt</label>
      <Input placeholder="Time!" />
    </div>
    <div>
      <label>FieldType</label>
      <Input placeholder="ID!" />
    </div>
    <div>
      <label>StringValue</label>
      <Input placeholder="String" />
    </div>
    <div>
      <label>IntValue</label>
      <Input placeholder="Int" />
    </div>
    <div>
      <label>FloatValue</label>
      <Input placeholder="Float" />
    </div>
    <div>
      <label>BooleanValue</label>
      <Input placeholder="Boolean" />
    </div>
    <div>
      <label>DateTimeValue</label>
      <Input placeholder="Time" />
    </div>
    <div>
      <label>Component</label>
      <Input placeholder="ID" />
    </div>

    <Button label="Edit" />
  </div>
)

CreateField.propTypes = {
  id: PropTypes.string,
}

export default CreateField
