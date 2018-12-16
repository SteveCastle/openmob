'use strict';

/**
 * Event.js controller
 *
 * @description: A set of functions called "actions" for managing `Event`.
 */

module.exports = {

  /**
   * Retrieve event records.
   *
   * @return {Object|Array}
   */

  find: async (ctx) => {
    if (ctx.query._q) {
      return strapi.services.event.search(ctx.query);
    } else {
      return strapi.services.event.fetchAll(ctx.query);
    }
  },

  /**
   * Retrieve a event record.
   *
   * @return {Object}
   */

  findOne: async (ctx) => {
    return strapi.services.event.fetch(ctx.params);
  },

  /**
   * Count event records.
   *
   * @return {Number}
   */

  count: async (ctx) => {
    return strapi.services.event.count(ctx.query);
  },

  /**
   * Create a/an event record.
   *
   * @return {Object}
   */

  create: async (ctx) => {
    return strapi.services.event.add(ctx.request.body);
  },

  /**
   * Update a/an event record.
   *
   * @return {Object}
   */

  update: async (ctx, next) => {
    return strapi.services.event.edit(ctx.params, ctx.request.body) ;
  },

  /**
   * Destroy a/an event record.
   *
   * @return {Object}
   */

  destroy: async (ctx, next) => {
    return strapi.services.event.remove(ctx.params);
  },

  /**
   * Add relation to a/an event record.
   *
   * @return {Object}
   */

  createRelation: async (ctx, next) => {
    return strapi.services.event.addRelation(ctx.params, ctx.request.body);
  },

  /**
   * Update relation to a/an event record.
   *
   * @return {Object}
   */

  updateRelation: async (ctx, next) => {
    return strapi.services.event.editRelation(ctx.params, ctx.request.body);
  },

  /**
   * Destroy relation to a/an event record.
   *
   * @return {Object}
   */

  destroyRelation: async (ctx, next) => {
    return strapi.services.event.removeRelation(ctx.params, ctx.request.body);
  }
};
