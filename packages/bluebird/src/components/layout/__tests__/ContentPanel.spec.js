import React from 'react';
import { shallow } from 'enzyme';
import ContentPanel from '../ContentPanel';

it('renders without crashing', () => {
  shallow(<ContentPanel />);
});
