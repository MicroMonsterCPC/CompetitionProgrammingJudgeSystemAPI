require "faraday"
require "json"

class QuestionAnswer
  def self.send_new_action(action, question_id, question_answer)
    data = {
      "action": action,
      "id": question_id.to_s,
      "answer": question_answer
    }
    client = Faraday.new(:url => "http://api:1323")
    res = client.post do |req|
      req.url '/create-answer'
      req.headers['Content-Type'] = 'application/json'
      req.body = data.to_json
    end
    return JSON.parse(res.body)["Result"]
  end
end

class QuestionsController < ApplicationController

  def index
    @questions = Question.all
  end

  def show
    @question = Question.find_by(id: params[:id])
  end

  def new
    @question = Question.new
  end

  def create
    @question = Question.new(question_params)

    respond_to do |format|
      if @question.save 
        if QuestionAnswer.send_new_action("create", @question.id, @question.answer)
          format.html { redirect_to @question, notice: 'Question was successfully created.' }
          format.json { render :show, status: :created, location: @question }
        else 
          @question.destroy if QuestionAnswer.send_new_action("Delete", @question.id, "")
        end
      else
        format.html { render :new }
        format.json { render json: @question.errors, status: :unprocessable_entity }
      end
    end
  end

  def update
    respond_to do |format|
      if @question.update(question_params) && QuestionAnswer.send_new_action("Update", @question.id, @question.answer)
        format.html { redirect_to @question, notice: 'Question was successfully updated.' }
        format.json { render :show, status: :ok, location: @question }
      else
        format.html { render :edit }
        format.json { render json: @question.errors, status: :unprocessable_entity }
      end
    end
  end

  def destroy
    @question.destroy
    respond_to do |format|
      format.html { redirect_to users_url, notice: 'Question was successfully destroyed.' }
      format.json { head :no_content }
    end
  end

  private

  def current_user
    return User.find_by(uid: session[:user_id])
  end
  # Never trust parameters from the scary internet, only allow the white list through.
  def question_params
    params.require(:question).permit(:title, :body, :answer)
  end

end
