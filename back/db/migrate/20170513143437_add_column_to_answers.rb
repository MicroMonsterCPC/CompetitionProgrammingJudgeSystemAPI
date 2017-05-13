class AddColumnToAnswers < ActiveRecord::Migration[5.0]
  def change
    add_column :answers, :lang, :string
  end
end
