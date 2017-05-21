class CreateAnswers < ActiveRecord::Migration[5.0]
  def change
    create_table :answers do |t|
      t.integer :question_id
      t.text :code
      t.string :lang

      t.timestamps
    end
  end
end
