import React from 'react';
import { shallow } from 'enzyme';
import Form from '../Form';

it('renders without crashing', () => {
  shallow(<Form />);
});
