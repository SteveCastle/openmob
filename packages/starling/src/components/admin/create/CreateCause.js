import React from 'react'
import PropTypes from 'prop-types'
import Input from '@openmob/bluebird/src/components/forms/Input'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const CreateCause = ({ id }) => (
  <div>
    <h1>Create Cause</h1>
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
      <label>Slug</label>
      <Input placeholder="String!" />
    </div>
    <div>
      <label>Summary</label>
      <Input placeholder="String" />
    </div>
    <div>
      <label>HomePage</label>
      <Input placeholder="ID" />
    </div>
    <div>
      <label>Photo</label>
      <Input placeholder="ID" />
    </div>

    <Button label="Edit" />
  </div>
)

CreateCause.propTypes = {
  id: PropTypes.string,
}

export default CreateCause
