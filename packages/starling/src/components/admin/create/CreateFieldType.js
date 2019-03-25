import React from 'react'
import PropTypes from 'prop-types'
import Input from '@openmob/bluebird/src/components/forms/Input'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const CreateFieldType = ({ id }) => (
  <div>
    <h1>Create FieldType</h1>
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
      <label>Title</label>
      <Input placeholder="String!" />
    </div>
    <div>
      <label>DataType</label>
      <Input placeholder="String!" />
    </div>
    <div>
      <label>PropName</label>
      <Input placeholder="String!" />
    </div>
    <div>
      <label>StringValueDefault</label>
      <Input placeholder="String" />
    </div>
    <div>
      <label>IntValueDefault</label>
      <Input placeholder="Int" />
    </div>
    <div>
      <label>FloatValueDefault</label>
      <Input placeholder="Float" />
    </div>
    <div>
      <label>BooleanValueDefault</label>
      <Input placeholder="Boolean" />
    </div>
    <div>
      <label>DateTimeValueDefault</label>
      <Input placeholder="Time" />
    </div>
    <div>
      <label>ComponentType</label>
      <Input placeholder="ID" />
    </div>

    <Button label="Edit" />
  </div>
)

CreateFieldType.propTypes = {
  id: PropTypes.string,
}

export default CreateFieldType
