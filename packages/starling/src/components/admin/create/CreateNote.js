import React from 'react'
import PropTypes from 'prop-types'
import Input from '@openmob/bluebird/src/components/forms/Input'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const CreateNote = ({ id }) => (
  <div>
    <h1>Create Note</h1>
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
      <label>Contact</label>
      <Input placeholder="ID!" />
    </div>
    <div>
      <label>Cause</label>
      <Input placeholder="ID!" />
    </div>
    <div>
      <label>Body</label>
      <Input placeholder="String" />
    </div>

    <Button label="Edit" />
  </div>
)

CreateNote.propTypes = {
  id: PropTypes.string,
}

export default CreateNote
