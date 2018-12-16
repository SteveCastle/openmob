'use strict';

/**
 * Product.js controller
 *
 * @description: A set of functions called "actions" for managing `Product`.
 */

module.exports = {

  /**
   * Retrieve product records.
   *
   * @return {Object|Array}
   */

  find: async (ctx) => {
    if (ctx.query._q) {
      return strapi.services.product.search(ctx.query);
    } else {
      return strapi.services.product.fetchAll(ctx.query);
    }
  },

  /**
   * Retrieve a product record.
   *
   * @return {Object}
   */

  findOne: async (ctx) => {
    return strapi.services.product.fetch(ctx.params);
  },

  /**
   * Count product records.
   *
   * @return {Number}
   */

  count: async (ctx) => {
    return strapi.services.product.count(ctx.query);
  },

  /**
   * Create a/an product record.
   *
   * @return {Object}
   */

  create: async (ctx) => {
    return strapi.services.product.add(ctx.request.body);
  },

  /**
   * Update a/an product record.
   *
   * @return {Object}
   */

  update: async (ctx, next) => {
    return strapi.services.product.edit(ctx.params, ctx.request.body) ;
  },

  /**
   * Destroy a/an product record.
   *
   * @return {Object}
   */

  destroy: async (ctx, next) => {
    return strapi.services.product.remove(ctx.params);
  },

  /**
   * Add relation to a/an product record.
   *
   * @return {Object}
   */

  createRelation: async (ctx, next) => {
    return strapi.services.product.addRelation(ctx.params, ctx.request.body);
  },

  /**
   * Update relation to a/an product record.
   *
   * @return {Object}
   */

  updateRelation: async (ctx, next) => {
    return strapi.services.product.editRelation(ctx.params, ctx.request.body);
  },

  /**
   * Destroy relation to a/an product record.
   *
   * @return {Object}
   */

  destroyRelation: async (ctx, next) => {
    return strapi.services.product.removeRelation(ctx.params, ctx.request.body);
  }
};
