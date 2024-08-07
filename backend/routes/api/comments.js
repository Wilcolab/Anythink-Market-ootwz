const router = require("express").Router();
const mongoose = require("mongoose");
const Comment = mongoose.model("Comment");

module.exports = router;

/**
 * Retrieves all comments.
 * 
 * @route GET /api/comments
 * @returns {Array} An array of comments.
 * @throws {Error} If there is an error retrieving the comments.
 */


router.get("/", (req, res) => {
  Comment.find()
    .then((comments) => {
      res.json(comments);
    })
    .catch((err) => {
      console.error(err);
      res.sendStatus(500);
    });
}
/**
 * Deletes a comment by ID.
 * 
 * @route DELETE /api/comments/:id
 * @param {string} id - The ID of the comment to delete.
 * @returns {number} HTTP status code 204 if the comment is deleted successfully.
 * @throws {Error} If there is an error deleting the comment.
 */

router.delete("/:id", async (req, res) => {
    try {
        await Comment.findByIdAndDelete(req.params.id);
        res.sendStatus(204);
    } catch (err) {
        console.error(err);
        res.sendStatus(500);
    }
});
