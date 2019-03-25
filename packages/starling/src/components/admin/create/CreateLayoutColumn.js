import React from 'react'
import PropTypes from 'prop-types'
import Input from '@openmob/bluebird/src/components/forms/Input'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const CreateLayoutColumn = ({ id }) => (
  <div>
    <h1>Create LayoutColumn</h1>
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
      <label>LayoutRow</label>
      <Input placeholder="ID!" />
    </div>
    <div>
      <label>Width</label>
      <Input placeholder="Int!" />
    </div>

    <Button label="Edit" />
  </div>
)

CreateLayoutColumn.propTypes = {
  id: PropTypes.string,
}

export default CreateLayoutColumn
