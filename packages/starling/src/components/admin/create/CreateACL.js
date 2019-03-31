import React from 'react'
import PropTypes from 'prop-types'
import Content from '@openmob/bluebird/src/components/layout/Content'
import Card from '@openmob/bluebird/src/components/cards/Card'
import Form from '@openmob/bluebird/src/components/forms/Form'
import Widget from '@openmob/bluebird/src/components/forms/Widget'
import Label from '@openmob/bluebird/src/components/forms/Label'
import Input from '@openmob/bluebird/src/components/forms/Input'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const CreateACL = ({ id }) => (
  <Content>
    <Card>
      <Form>
        <h1>Create ACL</h1>

        <Button label="Create" block variant="primary" />
      </Form>
    </Card>
  </Content>
)

CreateACL.propTypes = {
  id: PropTypes.string,
}

export default CreateACL
