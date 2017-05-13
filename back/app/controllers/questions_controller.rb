require "faraday"
require "json"

class Hoge
  def self.send_new_question(action, question_id, question_answer)
    puts action
    puts question_id
    puts question_answer
    data = {
      action: action,
      id: question_id,
      answer: question_answer
    }
    client = Faraday.new(:url => "http://api:1323")
    res = client.post("/create-answer", data)
    body = JSON.parse res.body
    p body
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
        Hoge.send_new_question("create", @question.id,@question.answer)
        format.html { redirect_to @question, notice: 'Question was successfully created.' }
        format.json { render :show, status: :created, location: @question }
      else
        format.html { render :new }
        format.json { render json: @question.errors, status: :unprocessable_entity }
      end
    end
  end

  def update
    respond_to do |format|
      if @user.update(question_params)
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
